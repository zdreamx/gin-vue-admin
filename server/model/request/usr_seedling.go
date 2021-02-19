package request

import "gin-vue-admin/model"

type SeedlingSearch struct{
    model.Seedling
    PageInfo
}