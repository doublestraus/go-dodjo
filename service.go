package dodjo

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Service struct {
	Client
}

func New(apiUrl, apiToken string) *Service {
	return &Service{
		Client{
			apiUrl:   apiUrl,
			apiToken: apiToken},
	}
}

func (service *Service) GetProducts() []*Product {
	pProducts := &Products{}
	err := service.makeRequest("GET", "/products/", []byte{}, false, nil, pProducts)
	if err != nil {
		logrus.Fatal(err)
	}
	productsData := make([]*Product, 0)
	for _, prod := range pProducts.Results {
		prod.Client = &service.Client
		productsData = append(productsData, &prod)
	}
	return productsData
}

func (service *Service) GetProduct(pid int) *Product {
	pProduct := &Product{}
	err := service.makeRequest("GET", fmt.Sprintf("/products/%d/", pid), []byte{}, false, nil, pProduct)
	if err != nil {
		logrus.Fatal(err)
	}
	pProduct.Client = &service.Client
	return pProduct
}

func (service *Service) GetProductByName(prodName string) []*Product {
	products := &Products{}
	err := service.makeRequest("GET", fmt.Sprintf("/products/?name=%s", prodName), []byte{}, false, nil, products)
	if err != nil {
		logrus.Fatal(err)
	}
	if products.Count == 0 {
		return []*Product{}
	}
	productsData := make([]*Product, 0)
	for _, prod := range products.Results {
		prod.Client = &service.Client
		productsData = append(productsData, &prod)
	}
	return productsData
}

func (service *Service) GetProductByNameOne(prodName string) *Product {
	products := service.GetProductByName(prodName)
	if len(products) != 0 {
		return products[0]
	} else {
		return nil
	}
}

func (service *Service) AddProduct(name, description string, productType int) *Product {
	product := &Product{Name: name, Description: description}
	product.EnableSimpleRiskAcceptance = true
	product.EnableFullRiskAcceptance = true
	product.ExternalAudience = true
	product.InternetAccessible = true
	product.ProdType = productType
	product.Created = time.Now()
	product.Tags = []string{"dd-api"}
	body, err := json.Marshal(product)
	if err != nil {
		logrus.Panic(err)
	}
	err = service.makeRequest("POST", "/products/", body, false, nil, product)
	if err != nil {
		logrus.Fatal(err)
	}
	product.Client = &service.Client
	return product
}
