package yantoml

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

type Yantoml struct {
	tree *toml.Tree
}

func New() *Yantoml {
	return &Yantoml{}
}

func (y *Yantoml) Read(file string) {
	var err error
	y.tree, err = toml.LoadFile(file)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
}

func (y *Yantoml) Add(file string, doc string) {
	y.Read(file)
	tree, _ := toml.Load(doc)
	y.tree = mergeTrees(y.tree, tree)
}

func mergeTrees(t1, t2 *toml.Tree) *toml.Tree {
	for _, key := range t2.Keys() {
		t1.Set(key, t2.Get(key))
	}
	return t1
}

func (y *Yantoml) Modify(file string, key string, value string) {
	y.Read(file)
	y.tree.Set(key, value)
}

func (y *Yantoml) Remove(file string, key string) {
	y.Read(file)
	y.tree.Delete(key)
}

func (y *Yantoml) Convert(file string) (string, error) {
	y.Read(file)
	tomlString, err := y.tree.ToTomlString()
	if err != nil {
		return "", err
	}
	return tomlString, nil
}
