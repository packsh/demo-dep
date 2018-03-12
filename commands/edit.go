package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/evanphx/json-patch"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util/editor"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	"encoding/json"
	api "github.com/kubepack/pack-server/apis/manifest/v1alpha1"
	packapi "github.com/kubepack/pack-server/apis/manifest/v1alpha1"
	ro_schema "k8s.io/apimachinery/pkg/runtime/schema"
)

const defaultEditor = "nano"
const _VendorFolder = "vendor"
const PatchFolder = "patch"

var (
	srcPath  string
	fileInfo os.FileInfo
)
// Local directory path needs to be absolute path. Patch filepath needs to be either absolute path or relative path.
func NewEditCommand(plugin bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit (filename)",
		Short: "Edit resource definition",
		Long:  "Generates patch via edit command",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			srcPath, err = cmd.Flags().GetString("patch")
			if err != nil {
				log.Println(err)
			}
			err = RunEdit(cmd, plugin)
			if err != nil {
				log.Println(err)
			}
		},
	}

	return cmd
}

func RunEdit(cmd *cobra.Command, plugin bool) error {
	root, err := cmd.Flags().GetString("file")
	if err != nil {
		return errors.WithStack(err)
	}
	if !plugin && !filepath.IsAbs(root) {
		wd, err := os.Getwd()
		if err != nil {
			return errors.WithStack(err)
		}
		root = filepath.Join(wd, root)
	}
	if !filepath.IsAbs(root) {
		return errors.Errorf("Duh! we need an absolute path when used as a kubectl plugin. For more info, see here: https://github.com/kubernetes/kubectl/issues/346")
	}
	path := filepath.Join(root, srcPath)
	if filepath.IsAbs(srcPath) {
		path = srcPath
	}
	fileInfo, err = os.Stat(path)
	if err != nil {
		return errors.WithStack(err)
	}

	srcFile, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.WithStack(err)
	}

	buf := &bytes.Buffer{}
	buf.Write(srcFile)

	edit := NewDefaultEditor()
	edited, _, err := edit.LaunchTempFile(fmt.Sprintf("%s-edit-", filepath.Base(os.Args[0])), ".yaml", buf)

	srcJson, err := yaml.YAMLToJSON(srcFile)
	if err != nil {
		return errors.WithStack(err)
	}

	dstJson, err := yaml.YAMLToJSON(edited)
	if err != nil {
		return errors.WithStack(err)
	}

	return GetPatch(srcJson, dstJson, cmd, plugin)
}

func GetPatch(src, dst []byte, cmd *cobra.Command, plugin bool) error {
	var err error
	var patch []byte

	// ref: https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/util/editor/editoptions.go#L549
	var ro runtime.TypeMeta
	if err := yaml.Unmarshal(src, &ro); err != nil {
		return errors.WithStack(err)
	}
	kind := ro.GetObjectKind().GroupVersionKind()
	versionedObject, err := scheme.Scheme.New(kind)

	switch {
	case runtime.IsNotRegisteredError(err):
		patch, err = jsonpatch.CreateMergePatch(src, dst)
	case err != nil:
		return errors.WithStack(err)
	default:
		patch, err = strategicpatch.CreateTwoWayMergePatch(src, dst, versionedObject)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	finalPatch, err := appendHeaderToPatch(src, patch)
	if err != nil {
		return errors.WithStack(err)
	}

	yamlPatch, err := yaml.JSONToYAML(finalPatch)
	if err != nil {
		return errors.WithStack(err)
	}
	root, err := cmd.Flags().GetString("file")
	if err != nil {
		return errors.WithStack(err)
	}
	if !plugin && !filepath.IsAbs(root) {
		wd, err := os.Getwd()
		if err != nil {
			return errors.WithStack(err)
		}
		root = filepath.Join(wd, root)
	}
	if !filepath.IsAbs(root) {
		return errors.Errorf("Duh! we need an absolute path when used as a kubectl plugin. For more info, see here: https://github.com/kubernetes/kubectl/issues/346")
	}
	patchFolderDir := strings.Replace(srcPath, _VendorFolder, PatchFolder, 1)

	repoDir := getRepositoryPath(patchFolderDir)
	dstPath := filepath.Join(root, api.ManifestDirectory, PatchFolder, repoDir)
	patchFileName, err := getPatchFileName(finalPatch)
	if err != nil {
		return errors.WithStack(err)
	}
	err = os.MkdirAll(dstPath, 0755)
	if err != nil {
		return errors.WithStack(err)
	}
	patchFilePath := filepath.Join(dstPath, patchFileName)
	_, err = os.Create(patchFilePath)
	if err != nil {
		return errors.WithStack(err)
	}
	err = ioutil.WriteFile(patchFilePath, yamlPatch, 0755)
	if err != nil {
		return errors.WithStack(err)
	}
	err = appendPatchToDependencies(filepath.Join(root, api.DependencyFile), patchFilePath)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// append apiVersion, kind and metadata.name with generated patch
func appendHeaderToPatch(src, patch []byte) ([]byte, error) {
	srcMap := map[string]interface{}{}
	err := json.Unmarshal(src, &srcMap)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	patchMap := map[string]interface{}{}
	err = json.Unmarshal(patch, &patchMap)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	patchMap["apiVersion"] = srcMap["apiVersion"]
	patchMap["kind"] = srcMap["kind"]
	patchMap["metadata"] = map[string]interface{}{}
	patchMap["metadata"].(map[string]interface{})["name"] = srcMap["metadata"].(map[string]interface{})["name"]

	finalPatch, err := json.Marshal(patchMap)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return finalPatch, nil
}

func NewDefaultEditor() editor.Editor {
	return editor.Editor{
		Args:  []string{defaultEditor},
		Shell: false,
	}
}

func appendPatchToDependencies(manifestPath, patchPath string) error {
	dep, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return errors.WithStack(err)
	}
	depList := api.DependencyList{}
	err = yaml.Unmarshal(dep, &depList)
	if err != nil {
		return errors.WithStack(err)
	}
	for key, val := range depList.Items {
		if strings.Contains(patchPath, val.Package) {
			depList.Items[key].Patches = append(val.Patches, patchPath)
			break
		}
	}
	depByte, err:= yaml.Marshal(depList)
	if err != nil {
		return errors.WithStack(err)
	}
	err = ioutil.WriteFile(manifestPath, depByte, 0755)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func getRepositoryPath(path string) string {
	splitFirst := strings.Split(path, filepath.Join(api.ManifestDirectory, PatchFolder))[1]
	spliFinal := strings.Split(splitFirst, filepath.Join(api.ManifestDirectory, "app"))[0]
	return strings.Trim(spliFinal, "/")
}

func getPatchFileName(patch []byte) (string, error) {
	patchStruct := &packapi.Dependency{}
	err := yaml.Unmarshal(patch, patchStruct)
	if err != nil {
		return "", errors.WithStack(err)
	}
	var ro runtime.TypeMeta
	if err := yaml.Unmarshal(patch, &ro); err != nil {
		return "", errors.WithStack(err)
	}
	gvk, err := ro_schema.ParseGroupVersion(patchStruct.APIVersion)
	if err != nil {
		return "", errors.WithStack(err)
	}
	name := strings.Join([]string{patchStruct.Name, patchStruct.Kind, gvk.Group, "yaml"}, ".")
	return name, nil
}
