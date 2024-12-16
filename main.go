package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Req struct {
	Expr string `json:"expression"`
}

type Res struct {
	Res string `json:"result,omitempty"`
	Err string `json:"error,omitempty"`
}

func Eval(exp interface{}) float64 {
	switch exp := exp.(type) {
	case *ast.ParenExpr:
		return Eval(exp.X)
	case *ast.BinaryExpr:
		return evalBinary(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.FLOAT, token.INT:
			i, _ := strconv.ParseFloat(exp.Value, 64)
			return i
		}
	}
	return 0
}

func evalBinary(exp *ast.BinaryExpr) float64 {
	l := Eval(exp.X)
	r := Eval(exp.Y)
	switch exp.Op {
	case token.ADD:
		return l + r
	case token.SUB:
		return l - r
	case token.MUL:
		return l * r
	case token.QUO:
		return l / r
	}
	return 0
}

func Calc(s string) (float64, error) {
	exp, err := parser.ParseExpr(s)
	if err != nil {
		return 0, err
	}
	return Eval(exp), nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp := Res{Err: "Internal server error"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer r.Body.Close()

	var req Req
	err = json.Unmarshal(body, &req)
	if err != nil || req.Expr == "" {
		resp := Res{Err: "Expression is not valid"}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}

	res, err := Calc(req.Expr)
	if err != nil || math.IsInf(res, 0) || math.IsNaN(res) {
		resp := Res{Err: "Expression is not valid"}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := Res{Res: fmt.Sprintf("%f", res)}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/v1/calculate", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
