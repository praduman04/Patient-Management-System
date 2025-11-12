package transformer

type UpdatePatient struct {
	Name   *string `json:"name" bson:"name,omitempty"`
	Email  *string `json:"email" bson:"email,omitempty"`
	Phone  *string `json:"phone" bson:"phone,omitempty"`
	Age    *int    `json:"age" bson:"age,omitempty"`
	Gender *string `json:"gender" bson:"gender,omitempty"`
}

//Using pointers (*string, *int) helps detect whether a field was provided or not.
//If nil, it means the field wasn’t sent in the JSON.
