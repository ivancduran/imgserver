package api

import (
	"os/exec"
	"strconv"
	"strings"
)

func FaceCMD(sp string) [][]int {
	cmd := "/usr/local/bin/facedetect " + sp
	f, _ := exec.Command("sh", "-c", cmd).Output()

	par := strings.Split(string(f), "\n")

	values := [][]int{}

	for _, v := range par {

		el := strings.Fields(v)

		ndata := []int{}

		for i := range el {
			nel, _ := strconv.Atoi(el[i])
			ndata = append(ndata, nel)
		}

		if len(ndata) > 0 {
			values = append(values, ndata)
		}

	}

	return values

}
