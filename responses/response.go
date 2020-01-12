package responses

import "time"

// GeneralResponse Respuesta generica de la app.
type GeneralResponse struct {
	Time string
	Data string
}

// SetDataResponse Establece los datos que seran encapsulados para la respuesta.
func (res* GeneralResponse) SetDataResponse(in string) {
	res.Data = in
}

// GetInitResponse Devuelve el contenido del objeto Data
func (res GeneralResponse) GetInitResponse() GeneralResponse {
	res.Time = time.Now().String()
	return res;
}
