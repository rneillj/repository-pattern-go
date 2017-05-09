package service

import(
    "log"

    "github.com/censhin/repository-pattern-go/repository"
    "github.com/censhin/repository-pattern-go/repository/db"
    "github.com/censhin/repository-pattern-go/service/domain"
)

func ListPokemon() ([]domain.Pokemon, err) {
    var data []domain.Pokemon

    ret, err := repository.List(&data, db.ListPokemonQuery)

    return ret, err
}

func GetPokemon() (domain.Pokemon, err) {
    var data domain.Pokemon

    ret, err := repository.Get(&data, db.PokemonQuery)

    return ret, err
}
