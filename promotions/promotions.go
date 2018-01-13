package promotions

import (
  "gopkg.in/yaml.v2"
  "log"
  "io/ioutil"
  "path/filepath"
)

type Promotions struct {
  Promotions []Promotion  `yaml:"promotions"`
}

type Promotion struct {
  PromoType  string `yaml:"type"`
  Name  string `yaml:"name"`
  Condition int `yaml:"condition"`
  Discount int `yaml:"discount"`
  GetFree int `yaml:"get_free"`
}

func (p *Promotions) GetPromotions() *Promotions {
  filename, _ := filepath.Abs("./config/promotions.yml")
  
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
