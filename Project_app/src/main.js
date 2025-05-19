import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from "element-plus"
import "element-plus/dist/index.css"

// 引入Vant组件库
import Vant from "vant";
import "vant/lib/index.css";

// 换算rem基准值
import "lib-flexible";

import {createRouter,createWebHistory} from "vue-router"

let router = createRouter({
  history:createWebHistory(),
  routes:[
    {
      path:"/",
      component:()=>import("./views/Home.vue")
    },
    {
      path:"/Detail/:CarId",
      component:()=>import("./views/Detail.vue")
    },
    {
      path  :"/NewCar/",
      component:()=>import("./views/NewCar.vue")
    },
    {
      path  :"/Assess/",
      component:()=>import("./views/Assess.vue")
    }
  ]

})

createApp(App).use(ElementPlus).use(Vant).use(router).mount('#app')
