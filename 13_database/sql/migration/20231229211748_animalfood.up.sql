create table if not exists animalFood (
    idAnimal uuid not null references animal(id),
    idFood uuid not null references food,

    primary key (idAnimal, idFood)
);