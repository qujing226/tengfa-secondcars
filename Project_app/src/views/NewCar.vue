<template>
  <div>
    <!--    头部  -->
    <!--    第一行信息-->
    <div class="header">
      <img
          src="data:image/png;base64,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"
          alt=" " >
    </div>
    <hr>
    <div class="btns">
      <!--        品牌下拉列表-->
      <el-button  class="commonElButton" @click="drawerButton = true"
      > 品牌
      </el-button>
      <!--   品牌下拉抽屉组件     -->
      <el-drawer
          v-model="drawerButton"
          size="85%"
          direction="btt"
          :before-close="handleClose"
      >
        <div>
          <el-card style="
              width: 700px;
              display: flex;
              align-items: center;
              margin-bottom: 10px;
              padding: 0;
            "
                   shadow="hover"
                   v-for="brand in brandList"
                   @click="drawerSeriesButton = true; getCarSeries(brand.id)"
          >
            <div style="display: flex; align-items: center;"> <!-- 保持这个div来管理图片和文本的对齐 -->
              <img :src="brand.imageUrl" alt="" style="width: 50px; margin-right: 20px;"/>
              <span style="font-size: 40px; text-align: center;">{{ brand.brand.split("-")[1] }}</span> <!-- 文本居中 -->
            </div>
          </el-card>
          <el-drawer
              v-model="drawerSeriesButton"
              size="60%"
              direction="rtl"
              :before-close="handleClose"

          >
            <div class="flex flex-wrap gap-4">
              <el-card
                  style="width: 400px;height:50px; margin-bottom: 6px;
                                display: flex; flex-direction: column; justify-content: center; align-items: center;"
                  @click="tipToHomePage(series.id)"
                  shadow="hover"
                  v-for="series in seriesList"
              >
                <!--                            p标签居中-->

                <p style="text-align: left;">{{ series.series }}</p>
              </el-card>
            </div>

          </el-drawer>
        </div>
      </el-drawer>

      <!--        搜索框-->
      <el-input v-model="keyword" type="text"

                placeholder="          Please input"></el-input>
      <el-button  class="commonElButton"@click="search">搜索
      </el-button>

      <!--        -->


      <el-select class="custom-select"
                 v-model="modelSearch"
                 placeholder="Select"
                 size="large"
                 @change="getCarData"

      >
        <el-option
            v-for="model in models"
            :label="model.model"
            :value="model.id">
        </el-option>
      </el-select>


      <hr>
      <el-button
          class="four-button-group"
          type:plain
          v-for="(saleType,index) in saleType"
          color="#d2d2c8"
          :style="{color:saleType.id==saleTypeInt?'red':'black'}"
          @click="chooseCarsSaleType(saleType.id),saleTypeInt=saleType.id"
      >{{ saleType.saleType }}
      </el-button>
    </div>

    <!--    中部-->
    <ul class = "ulli"style="width:370px">
      <li >
        <span
            class="text-truncate" >奥迪A8 2023款 A8L 50 TFSI quattro 臻选型</span>
        <span
            :style="{
                display: 'inline-block',
                marginLeft: '16px',
                fontSize: '10px',
                padding: '6px',
                textAlign: 'center',
                borderRadius: '16px',
                height: '13px',
                width: '60px',
                color: 'white',
                letterSpacing: '3px',
                fontFamily: 'Times New Roman',
                backgroundColor: 'blue',
              }"
        >中规规车</span>
        <h4 class = "h4">70.99万</h4>
      </li>
      <li >
        <span
            class="text-truncate" >奥迪A8 2023款 A8L 50 TFSI quattro 臻选型</span>
        <span
            :style="{
                display: 'inline-block',
                marginLeft: '16px',
                fontSize: '10px',
                padding: '6px',
                textAlign: 'center',
                borderRadius: '16px',
                height: '13px',
                width: '60px',
                color: 'white',
                letterSpacing: '3px',
                fontFamily: 'Times New Roman',
                backgroundColor: 'blue',
              }"
        >中规规车</span>
        <h4  class = "h4">70.99万</h4>
      </li>
    </ul>



    <div class="empty"></div>
    <ul class="tabbar" >
      <li
          v-for="(item,index) in tabbar"
          :style="{color:index==tabbarIndex?'#2196f3':'black'}"
          @click="tabbarIndex=index,TipToChangPage(index)">
        <p class="iconfont icon" v-html="item.icon"></p>
        <p class="title">{{ item.title }}</p>
      </li>
    </ul>
  </div>
</template>
<script setup>
import FilterBar from './FilterBar.vue';

</script>
<script>

export default {
  components:{
    FilterBar
  },
  data() {
    return {
      tabbar: [
        {
          title: "首页",
          icon: "&#xe61d;",
        },
        {
          title: "二手车",
          icon: "&#xfd72;",
        },
        {
          title: "新车",
          icon: "&#xe610;",
        },
        {
          title: "卖车",
          icon: "&#xe6e9;",
        },
      ],
      tabbarIndex: 2,
      saleType: [{
        id: 0,
        saleType: "默认排序",
      }, {
        id: 1,
        saleType: "最新发布",
      }, {
        id: 2,
        saleType: "价格最高",
      }, {
        id: 3,
        saleType: "价格最低",
      }
      ],
      saleTypeInt: 100,
      drawerButton: false,
      drawerSeriesButton: false,
      brandList:  [
        {
          "id": 1,
          "brand": "A - 奥迪",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539288541462528.png"
        },
        {
          "id": 2,
          "brand": "A - 阿斯顿·马丁",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539179229511680.png"
        },
        {
          "id": 3,
          "brand": "A - 阿尔法·罗密欧",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539236502732800.png"
        },
        {
          "id": 4,
          "brand": "A - AC SCHNITZER",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266504955152564224.png"
        },
        {
          "id": 536,
          "brand": "A - ALPINA",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266596173756760064.png"
        },
        {
          "id": 20,
          "brand": "B - 保斐利",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266593922334064640.png"
        },
        {
          "id": 115,
          "brand": "B - BEIJING汽车",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538727825932288.png"
        },
        {
          "id": 156,
          "brand": "B - 北汽幻速",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538674138841088.png"
        },
        {
          "id": 167,
          "brand": "B - 北汽新能源",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538621508714496.png"
        },
        {
          "id": 172,
          "brand": "B - 宝沃",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537771625283584.png"
        },
        {
          "id": 499,
          "brand": "B - 比速",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266596467148324864.png"
        },
        {
          "id": 573,
          "brand": "B - 北汽道达",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602857183051776.png"
        },
        {
          "id": 654,
          "brand": "B - 铂驰",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140536818243207168.png"
        },
        {
          "id": 853,
          "brand": "B - 比德文汽车",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266519240440283136.png"
        },
        {
          "id": 2398,
          "brand": "B - 北汽瑞翔",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266495176497692672.png"
        },
        {
          "id": 2432,
          "brand": "B - 百智新能源",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266494408675819520.png"
        },
        {
          "id": 21,
          "brand": "C - 长安",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535818543431680.png"
        },
        {
          "id": 22,
          "brand": "C - 长城",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140536152246452224.png"
        },
        {
          "id": 23,
          "brand": "C - 长安欧尚",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535934515937280.png"
        },
        {
          "id": 24,
          "brand": "C - 昌河",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535874256371712.png"
        },
        {
          "id": 497,
          "brand": "C - 成功",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140536078917435392.png"
        },
        {
          "id": 1716,
          "brand": "C - 车驰汽车",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266512082277498880.png"
        },
        {
          "id": 1768,
          "brand": "C - 创维汽车",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266507189466693632.png"
        },
        {
          "id": 25,
          "brand": "D - 大众",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535757684080640.png"
        },
        {
          "id": 26,
          "brand": "D - 东风风行",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535304288206848.png"
        },
        {
          "id": 27,
          "brand": "D - 东南",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535585084276736.png"
        },
        {
          "id": 28,
          "brand": "D - 东风小康",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535304288206848.png"
        },
        {
          "id": 29,
          "brand": "D - 道奇",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535648351158272.png"
        },
        {
          "id": 30,
          "brand": "D - 东风风神",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
        },
        {
          "id": 31,
          "brand": "D - DS",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535693653835776.png"
        },
        {
          "id": 32,
          "brand": "D - 东风风度",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
        },
        {
          "id": 33,
          "brand": "D - 东风",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
        },
        {
          "id": 170,
          "brand": "D - 东风风光",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
        },
        {
          "id": 574,
          "brand": "D - 电咖",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602661841731584.png"
        },
        {
          "id": 662,
          "brand": "D - 达契亚",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266526104389419008.png"
        },
        {
          "id": 1705,
          "brand": "D - 大运",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266514695194673152.png"
        },
        {
          "id": 2399,
          "brand": "D - 电动屋",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266494826730487808.png"
        },
        {
          "id": 35,
          "brand": "F - 福特",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534585615187968.png"
        },
        {
          "id": 36,
          "brand": "F - 丰田",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534705870077952.png"
        },
        {
          "id": 37,
          "brand": "F - 菲亚特",
          "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534529088552960.png"
        },
      ],
    };
  },
  methods:{
    TipToChangPage(ind){
      if (ind==0){
        this.$router.push(`/`)
      }else if (ind==3){
        this.$router.push(`/Assess/`)
      }
    },
  }


}
</script>
<style scoped>
* {
  margin: 0;
  padding: 0;
  font-family: Arial, sans-serif;
}
.header div{
  width: 750px;
}
.header img{
  height: 16px;width: 40px;margin: 8px
}
.four-button-group{
  margin-left:25px;
  width:70px;
  height: 20px;
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 11px;
  background-color: #cbcede;
}

hr{
  width:750px;
}
.commonElButton {
  margin-left: 8px;
  width: 50px;
  font-size: 17px;
  height: 30px;
}
.el-input{
  width: 235px;margin:16px;height: 30px;font-size: 17px
}

.text-truncate{
  font-size: 17px; padding:10px; margin-left:40px;
  font-family: 'Microsoft YaHei', sans-serif; font-weight: bold;
}

/*对tabbar进行样式调整*/
@font-face {
  font-family: "iconfont";
  src: url("../assets/iconfont/iconfont.woff2?t=1719390343114") format("woff2"),
  url("../assets/iconfont/iconfont.woff?t=1719390343114") format("woff"),
  url("../assets/iconfont/iconfont.ttf?t=1719390343114") format("truetype");
}

.iconfont {
  font-family: "iconfont" !important;
  font-size: 10px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.tabbar {
  display: flex;
  justify-content: space-around;
  background: #f2f2f2;
  padding: 0;
  margin:0;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  /* 定位 */
  position: fixed;
  left: 0;
  bottom: 0;
  width: 400px;
  height:60px;
}

.tabbar li {
  transition: all 0.5s;
  width: 40px;
}

.icon {
  text-align: center;width: 40px;
  font-size:25px;
  margin-top:10px;
  margin-bottom:0;
}

.title {
  text-align: center;
  width:40px;
  font-size:12px;
  margin-top:5px;
}
.ulli li{
  margin:20px;
  width:360px;
}
.h4{
  font-size:20px;color:red;margin-left: 16px;margin-top:4px
}
</style>