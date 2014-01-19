package vect 

import (
	"encoding/json"
)

func (v V) MarshalJSON() ([]byte, error) {
	return json.Marshal(&[2]float64{v.X, v.Y})
}

func (v *V) UnmarshalJSON(data []byte) error {
	vectData := [2]float64{}

	//try unmarshalling array form
	err := json.Unmarshal(data, &vectData)
	if err != nil {
		//try other form
		vectData := struct {
			X, Y float64
		}{}

		err := json.Unmarshal(data, &vectData)

		if err != nil {
			return err
		}
		v.X = vectData.X
		v.Y = vectData.Y
		return nil
	}

	v.X = vectData[0]
	v.Y = vectData[1]

	return nil
}
