package service

import(
    "log"

    "github.com/censhin/repository-pattern-go/repository"
    "github.com/censhin/repository-pattern-go/repository/db"
    "github.com/censhin/repository-pattern-go/service/domain"
)

func ListPokemon() ([]domain.Pokemon, err) {
    var data []domain.Pokemon
    err := repository.List(&data, db.ListPokemonQuery)
    return data, err
}

func GetPokemon(ID string) (domain.Pokemon, err) {
    var data domain.Pokemon
    err := repository.Get(ID, &data, db.PokemonQuery)
    return data, err
}
