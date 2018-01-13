package products

import (
  "gopkg.in/yaml.v2"
  "log"
  "io/ioutil"
  "path/filepath"
)

type Products struct {
  Products []Product  `yaml:"products"`
}

type Product struct {
  Code  int `yaml:"code"`
  Name  string `yaml:"name"`
  Price int `yaml:"price"`
}

func (p *Products) GetProducts() *Products {
  filename, _ := filepath.Abs("./config/products.yml")
  
  yamlFile, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Fatalf("error: %v", err)
  }

  err = yaml.Unmarshal(yamlFile, &p)
  if err != nil {
      panic(err)
  }
  return p
}
