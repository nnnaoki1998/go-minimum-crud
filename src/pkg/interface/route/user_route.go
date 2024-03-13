package route

import (
	"encoding/json"
	"go-minimum-crud/src/pkg/application/service"
	"go-minimum-crud/src/pkg/domain/model"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

type UserRoute struct {
	UserService service.UserService
}

func (c *UserRoute) index(w http.ResponseWriter, _ *http.Request) (string, model.StatusCode) {
	users, err := c.UserService.Index()
	if err != nil {
		log.Error().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusInternalServerError)
	}
	json, err := json.Marshal(users)
	if err != nil {
		log.Error().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusInternalServerError)
	}
	return string(json), model.StatusCode(http.StatusOK)
}

func (c *UserRoute) show(w http.ResponseWriter, r *http.Request) (string, model.StatusCode) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Warn().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusBadRequest)
	}
	user, err := c.UserService.Show(model.UserId(id))
	if err != nil {
		log.Error().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusInternalServerError)
	}
	json, err := json.Marshal(user)
	if err != nil {
		log.Error().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusInternalServerError)
	}
	return string(json), model.StatusCode(http.StatusOK)
}

func (c *UserRoute) create(w http.ResponseWriter, r *http.Request) (string, model.StatusCode) {
	var user model.NewUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Warn().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusBadRequest)
	}
	if err := c.UserService.Create(user); err != nil {
		log.Error().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusInternalServerError)
	}
	return "", model.StatusCode(http.StatusOK)
}

func (c *UserRoute) update(w http.ResponseWriter, r *http.Request) (string, model.StatusCode) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Warn().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusBadRequest)
	}
	var user model.NewUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Warn().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusBadRequest)
	}
	if err := c.UserService.Update(model.UserId(id), user); err != nil {
		log.Error().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusInternalServerError)
	}
	return "", model.StatusCode(http.StatusOK)
}

func (c *UserRoute) delete(w http.ResponseWriter, r *http.Request) (string, model.StatusCode) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Warn().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusBadRequest)
	}
	if err := c.UserService.Delete(model.UserId(id)); err != nil {
		log.Error().Msg(err.Error())
		return err.Error(), model.StatusCode(http.StatusInternalServerError)
	}
	return "", model.StatusCode(http.StatusOK)
}

func (c *UserRoute) SetRouter() {
	usersRoute := func(w http.ResponseWriter, r *http.Request) {
		content, statusCode := func() (string, model.StatusCode) {
			switch r.Method {
			case http.MethodGet:
				return c.index(w, r)
			default:
				return "", model.StatusCode(http.StatusNotFound)
			}
		}()
		if int(statusCode) == http.StatusOK {
			http.Error(w, content, int(statusCode))
		} else {
			w.WriteHeader(int(statusCode))
		}
	}

	userRoute := func(w http.ResponseWriter, r *http.Request) {
		content, statusCode := func() (string, model.StatusCode) {
			switch r.Method {
			case http.MethodGet:
				return c.show(w, r)
			case http.MethodPost:
				return c.create(w, r)
			case http.MethodPatch:
				return c.update(w, r)
			case http.MethodDelete:
				return c.delete(w, r)
			default:
				return "", model.StatusCode(http.StatusNotFound)
			}
		}()
		if int(statusCode) == http.StatusOK {
			http.Error(w, content, int(statusCode))
		} else {
			w.WriteHeader(int(statusCode))
		}
	}

	http.HandleFunc("/users", usersRoute)
	http.HandleFunc("/user", userRoute)
}
