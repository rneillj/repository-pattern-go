package service

type Service interface {
    ListPokemon func () ([]model.Pokemon, err)
    GetPokemon  func () (model.Pokemon, err)
}
