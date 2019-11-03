package route

import (
    _ "fmt"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/fat-max/pyro-api/model"
)

func AllChemicals(w http.ResponseWriter, r *http.Request) {
    chemicals := model.GetAllChemicals()
    w.Header().Set("Content-Type", "application/json")

    if data, err := json.Marshal(chemicals); err == nil {
        w.WriteHeader(http.StatusOK)
        w.Write(data)
        return
    }

    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(`{"error": "error marshalling data"}`))
}


func Chemical(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")

    if id, ok := pathParams["id"]; ok {
        data := model.GetChemical(id)

        if chem, err := json.Marshal(data); err == nil {
            w.WriteHeader(http.StatusOK)
            w.Write(chem)
            return
        }

        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"error": "id not found"}`))
        return
    }

    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"error": "missing id"}`))
}
