# formation-go

Ce service centralise les differents référentiels en provenance de PWA
Il existe 2 routes gérant les 2 types de référentiels: 
Pour chacune des routes, les codes pris en compte sont fournis dans le swagger.

Aussi, chacune des 2 routes acceptent le parametre de recherche ```?search=[search]``` permettant de filtrer les résultats via leur libellé. Le filtre s'effectue sur un ```contains``` du search dans le ```label``` des résultats remontés (case insensitive & unaccent)

## referentiel/{code} 
Cette route permet de remonter les référentiels "standard" de PWA (ex: countries)

## parameter/{code}:
Cette route permet de remonter les référentiels spécifiques au paramétrage Repam chez PWA (ex: professional_situations).

Pour chacun de ces référentiel, il est possible d'ajouter des métadonnées provenant du endpoint PWA: "variable-globale-historisee".

Ces ajouts se font de manière automatique en se basant sur le retour de la méthode ```GetParameterMetadata (lib/types/pwa_metadata.go)``` qui retourne pour chaque ```PARAMETER_KEY``` la liste des objets ```PWA_METADATA``` associées.

ex
```go
return map[PWA_PARAMETER_KEY][]PWAParameterMetadata{
    PROFESSIONNAL_SITUATION: {  // Clé du paramètre auquel associé les metadata         
        {
            Name: "classe",  // nom de la metadata à remonter
            Code: PWA_PROFESSION_CLASS,  // correspondance variable PWA pour récupérer la valeur associée
        },
        ...,
    },
}
```

## Exemples d'appel:

* **api//referentiel/countries?search=FRA**

Retour: 
```json
[
    {
        "value": "FR",
        "label": "France",
        "metadata": null
    },
    {
        "value": "GF",
        "label": "Guyane Française",
        "metadata": null
    },
    {
        "value": "PF",
        "label": "Polynésie Française",
        "metadata": null
    },
    {
        "value": "TF",
        "label": "Terres Australes Françaises",
        "metadata": null
    }
]
```

* **/parameter/professionnal_situations?search=Infirmier**

Retour:
```json
[
    {
        "value": "303",
        "label": "Infirmier",
        "metadata": {
            "classe": "3",
            "regime": "5"
        }
    }
]
```


## Generate swagger documentation 

https://github.com/swaggo/echo-swagger

Add swag command : 

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate folder `/docs` : 

```shell
swag init
```

## Run project
```shell
sh ./run.sh
```


## Work with live reload 

```shell
go install github.com/cosmtrek/air@latest
```

```shell
air main.go
```
