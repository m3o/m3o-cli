package cli

type IndexResponse struct {
	Apis []*ExploreAPI `protobuf:"bytes,1,rep,name=apis,proto3" json:"apis,omitempty"`
}

type SearchResponse struct {
	Apis []*ExploreAPI `protobuf:"bytes,1,rep,name=apis,proto3" json:"apis,omitempty"`
}

type ExploreAPI struct {
	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Category    string      `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Icon        string      `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Endpoints   []*Endpoint `protobuf:"bytes,5,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	DisplayName string      `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

type Endpoint struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

type PublicAPI struct {
	Id           string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string           `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	OpenApiJson  string           `protobuf:"bytes,4,opt,name=open_api_json,json=openApiJson,proto3" json:"open_api_json,omitempty"`
	Pricing      map[string]int64 `protobuf:"bytes,6,rep,name=pricing,proto3" json:"pricing,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // map of endpoints to price. Unit is 1/10,000ths of a cent which allows us to price in fractions e.g. a request costs 0.0001 cents or 10000 requests for 1 cent
	OwnerId      string           `protobuf:"bytes,7,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	ExamplesJson string           `protobuf:"bytes,8,opt,name=examples_json,json=examplesJson,proto3" json:"examples_json,omitempty"`
	Category     string           `protobuf:"bytes,9,opt,name=category,proto3" json:"category,omitempty"`
	Icon         string           `protobuf:"bytes,10,opt,name=icon,proto3" json:"icon,omitempty"`
	PostmanJson  string           `protobuf:"bytes,11,opt,name=postman_json,json=postmanJson,proto3" json:"postman_json,omitempty"`
	DisplayName  string           `protobuf:"bytes,12,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Quotas       map[string]int64 `protobuf:"bytes,13,rep,name=quotas,proto3" json:"quotas,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // map of endpoints to monthly free request quota. After quota is exhausted endpoint becomes paid
}

type ListResponse struct {
	Apis []*PublicAPI `protobuf:"bytes,1,rep,name=apis,proto3" json:"apis,omitempty"`
}
