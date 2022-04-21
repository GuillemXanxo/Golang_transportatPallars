package models

type Viatge struct {
  Origen string `json:"origen"`
  Desti string `json:"desti"`
  Places int `json:"places"`
  DataNumber int `json:"dataNumber"`
  Data string `json:"data"`
  HoraSortidaNumber int `json:"horaSortidaNumber"`
  HoraSortida string `json:"horaSortida"`
  Comentaris string `json:"comentaris"`
  Dones bool `json:"dones"`
}
