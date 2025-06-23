package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type OpenAPI struct {
	OpenAPI    string                   `yaml:"openapi"`
	Info       map[string]interface{}   `yaml:"info"`
	Servers    []map[string]interface{} `yaml:"servers,omitempty"`
	Paths      map[string]interface{}   `yaml:"paths"`
	Components map[string]interface{}   `yaml:"components,omitempty"`
}

func readYAMLFile(path string) (*OpenAPI, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var api OpenAPI
	err = yaml.Unmarshal(data, &api)
	if err != nil {
		return nil, err
	}

	return &api, nil
}

func mergeAPIs(apis []*OpenAPI) *OpenAPI {
	merged := &OpenAPI{
		OpenAPI:    "3.0.0",
		Info:       map[string]interface{}{"title": "Merged API", "version": "1.0.0"},
		Paths:      make(map[string]interface{}),
		Components: make(map[string]interface{}),
	}

	for _, api := range apis {
		for path, item := range api.Paths {
			merged.Paths[path] = item
		}

		if api.Components != nil {
			for key, comp := range api.Components {
				if merged.Components[key] == nil {
					merged.Components[key] = comp
				} else {
					// 可选：合并具体组件类型，如 schemas、responses 等
					if compMap, ok := comp.(map[string]interface{}); ok {
						existingMap := merged.Components[key].(map[string]interface{})
						for k, v := range compMap {
							existingMap[k] = v
						}
					}
				}
			}
		}
	}

	return merged
}

func writeYAMLFile(api *OpenAPI, path string) error {
	data, err := yaml.Marshal(api)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func GenApi() {
	// 遍历当前目录及子目录下的所有 openapi-*.yml 文件
	inputFiles, err := filepath.Glob("../api/openapi-*.yml")
	if err != nil {
		panic(err)
	}

	var apis []*OpenAPI
	for _, file := range inputFiles {
		println("openapi file:", file)
		api, err := readYAMLFile(file)
		if err != nil {
			fmt.Printf("读取失败: %s\n", file)
			continue
		}
		apis = append(apis, api)
	}

	merged := mergeAPIs(apis)
	mergedFile := "../api/merged-api.yaml"
	errMerge := writeYAMLFile(merged, mergedFile)
	if errMerge != nil {
		fmt.Println("写入合并文件失败:", err)
		return
	}

	fmt.Println(fmt.Sprintf("✅ 合并完成: %s", mergedFile))
	genCode(mergedFile)
	errDel := os.Remove(mergedFile)
	if errDel != nil {
		return
	}
}

func genCode(genFile string) {
	cmd := exec.Command("go", "run", "github.com/ogen-go/ogen/cmd/ogen@latest",
		"-config", "./gen-api-config.yml",
		"--target", "../gen/oas",
		"--package", "oas",
		"--clean", genFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("开始执行 ogen 命令...")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("执行失败: %v\n", err)
		return
	}

	fmt.Println("✅ ogen 执行完成")
}
