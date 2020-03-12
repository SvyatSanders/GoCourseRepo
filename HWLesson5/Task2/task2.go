package main

import (
	"fmt"
	"os"
)

// 2.	Что бы вы изменили в программе чтения из файла? Приведите исправленный вариант, обоснуйте свое решение в комментарии.
// Ответ. Если рассматривать первоначальный способ чтения файла (не упрощенный), то мне бросилось в глаза 3 одинаковые проверки на ошибку
// Сделал данную проверку в виде функции errr()
// Полагаю, что возможно сделать интерфейс. Попробовал сделать интерфейс для открытия файла.
// Уверен, что это кривое решение. Буду рад корректировкам что следовало сделать лучше.

type files interface {
	OpenFile()
}

type filePath struct {
	Path string
}

func (f filePath) OpenFile() *os.File {
	p, _ := os.Open(f.Path)
	return p
}

func errCheck(errors error) error {
	if errors != nil {
		return errors
	}
	return nil
}

func main() {
	x := filePath{Path: "/Users/SvyatSanders/go/src/GoCourseRepo/HWLesson5/Task2/testfile.txt"}
	fmt.Println(x.OpenFile())
	defer x.OpenFile().Close()

	// getting size of file
	stat, err := x.OpenFile().Stat()
	errCheck(err)

	// reading file
	bs := make([]byte, stat.Size())
	_, err = x.OpenFile().Read(bs)
	errCheck(err)

	fmt.Println(string(bs))

}
