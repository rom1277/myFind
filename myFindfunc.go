package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func myFind(path string, d, f, sl bool, ext string) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, file := range files {
		typefile := file.Mode()
		fullPath := filepath.Join(path, file.Name())
		if file.IsDir() {
			if d {
				fmt.Printf("Dir: %s\n", fullPath)

			}
			if err = myFind(fullPath, d, f, sl, ext); err != nil {
				return err
			}
		}
		if f {
			if ext != "" {
				fileExt := filepath.Ext(file.Name()) // Получаем расширение файла
				if fileExt == ext {                  // Сравниваем расширение
					fmt.Printf("File: %s\n", fullPath)
				}
			} else {
				if typefile&os.ModeType == 0 {
					fmt.Printf("File: %s\n", fullPath)
				}

			}
		}

		if sl {
			if typefile&os.ModeSymlink != 0 {
				ref, err := os.Readlink(fullPath)
				if err != nil {
					fmt.Printf("Link: %s -> [broken]\n", fullPath)
					continue
				}
				fullref := filepath.Join(filepath.Dir(fullPath), ref)
				_, err = os.Stat(fullref)
				if err != nil {
					fmt.Printf("Link: %s -> [broken]\n", fullPath)
					continue
				}
				fmt.Printf("Link: %s -> %s\n", fullPath, fullref)
			}
		}
	}
	return nil
}
