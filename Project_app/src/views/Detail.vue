<template>
  <div id="topDiv">
    <van-nav-bar
        class="fixed-navbar"
        title="车辆详情"
        left-text="返回"
        left-arrow
        @click-left="onClickLeft"

    />
    <div class="fixed-navdiv"></div>

    <!--  轮播图-->
    <van-swipe class="img-swipe" :autoplay="3000" indicator-color="#135eac">
      <van-swipe-item v-for="item in data.PhotoVo.Detail.split(',')">
        <img :src="item">
      </van-swipe-item>
    </van-swipe>
    <div class="bluePatten">
      <span class='span1'>{{ this.data.CurrentPrice }}</span>
      <span class="span2"> 万</span>
      <span class="span3">包含过户税</span>
      <span
          class="span4">新车完税价： {{ this.data.PriceVo.CarDownPay + this.data.PriceVo.CaHallPrice }}万  为您节约 {{ this.data.PriceVo.CarDownPay }}万</span>
    </div>
  </div>


  <div id="medium">
      <span class="saleType"
            v-if="data.SaleType !== undefined && saleType[data.SaleType]"
            :style="{backgroundColor: saleTypeColor(this.data.SaleType)}"
      > {{ saleType[this.data.SaleType].saleType }}</span>
    <h3 class="carname">{{ this.data.CarName }}</h3>
    <van-divider content-position="left" style="height:6px"
                 :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>
    <div class="MonthlyPay">
      <img src="../assets/按揭付款.png">
      <span> {{this.data.PriceVo.MonthlyRepayment}}元 ({{this.data.PriceVo.MortgageNum}})月</span>
    </div>
    <van-divider content-position="left" style="height:6px"
                 :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>
    <img class="tupian1" src="../assets/服务承诺.png"/>
    <img class="tupian2" src="https://tfcar.oss-cn-hangzhou.aliyuncs.com/513058784461979648.jpg"/>
  </div>

  <van-divider content-position="left" style="height:6px"
               :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>

  <span style="background-color:#204fbf;width:30px;margin-right:10px;padding-top:3px">&emsp;</span><span class="tag">基本信息</span>
  <van-divider content-position="left" style="height:6px"
               :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>
  <div class="basicInformation">
    <div class="list_1">
      <ul class="list_ul">
        <li class="FatherSpan">上牌时间&emsp; <span
            class='childrenSpan'>{{ this.data.DateOfRegistration.split("T")[0] }}</span></li>
        <li class="FatherSpan">车身颜色&emsp; <span class='childrenSpan'>{{ this.data.ParameterVo.CarColor }}</span>
        </li>
        <li class="FatherSpan">排量&emsp;&emsp;&emsp; <span
            class='childrenSpan'>{{ this.data.ParameterVo.EmissionStandard }}</span></li>
        <li class="FatherSpan">燃烧方式&emsp; <span class='childrenSpan'>{{ this.data.ParameterVo.Fuel }}</span></li>
        <li class="FatherSpan">出厂日期&emsp; <span class='childrenSpan'>—</span></li>
        <li class="FatherSpan">座位数&emsp;&emsp; <span class='childrenSpan'>{{ this.data.Seat }}</span></li>
        <li class="FatherSpan">车身描述&emsp; <span class='childrenSpan'>暂无</span></li>
      </ul>
    </div>
    <div class="list_1">
      <ul class="list_ul">
        <li class="FatherSpan">公里数&emsp;&emsp;<span class='childrenSpan'>{{ this.data.Mileage }}</span></li>
        <li class="FatherSpan">变速箱&emsp;&emsp;<span class='childrenSpan'>{{ this.data.ParameterVo.Gearbox }}</span>
        </li>
        <li class="FatherSpan">排放标准&emsp; <span class='childrenSpan'>{{ this.data.FuelType }}</span></li>
        <li class="FatherSpan">过户次数&emsp; <span class='childrenSpan'>—</span></li>
        <li class="FatherSpan">车架号&emsp;&emsp; <span class='childrenSpan'>{{ this.data.CarId }}</span></li>
        <li class="FatherSpan">驱动方式&emsp; <span class='childrenSpan'>{{ this.data.ParameterVo.Drive }}</span></li>
      </ul>
    </div>
  </div>
  <van-divider content-position="left" style="height:6px"
               :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>
  <span style="background-color:#204fbf;width:30px;margin-right:10px;padding-top:3px">&emsp;</span><span class="tag">详情图片</span>
  <van-divider content-position="left" style="height:6px"
               :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>
  <div id="carDetailPhotos">
    <ul v-for="(img,index) in this.data.PhotoVo.Detail.split(',')">
      <img :src="img" v-show="index <= visibleImageCount"/>
    </ul>
    <h4 style="color:grey;text-align: center;">注：车辆具体配置信息以实车为主</h4>
    <p @click="visibleImageCount=100" style="text-align: center;margin:20px;font-weight: bolder;">点击查看更多</p>
  </div>


  <img src="../assets/详情页底部.png" style="width:100%;"/>


</template>
<script>
import axios from "axios";

export default {
  name: "Detail",
  created() {
    // 1.保存传入的iid
    this.ind = this.$route.params.CarId;
    this.getDetail(this.ind)
  },
  data() {
    return {
      data: [],
      ind: "202112040957088478",
      visibleImageCount: 5,
      saleType: [{
        id: 0,
        saleType: "腾发自营",
      }, {
        id: 1,
        saleType: "车主直营",
      }, {
        id: 2,
        saleType: "限时特惠",
      }, {
        id: 3,
        saleType: "严选好车",
      }
      ],
    };
  },
  methods: {
    getDetail(CarId) {
      axios
          .get(`http://118.178.120.11:8080/detailInformation?carId=` + CarId, {})
          .then((res) => {
            // 轮播图
            console.log(res.data.carDetail.PhotoVo.Detail.split(","))
            this.data = res.data.carDetail;
          })
    },
    onClickLeft() {
      this.$router.push("/")
    }

  },
  computed: {
    saleTypeColor() {
      return (saleType) => {
        switch (saleType) {
          case 0:
            return 'orange';
          case 1:
            return 'blue';
          case 2:
            return 'green';
          default:
            return 'yellow';
        }
      };
    },
  },
}
</script>
<style scoped>
* {
  margin: 0;
  padding: 0;
}

.fixed-navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 999; /* 可以根据需要调整z-index，确保它高于其他元素 */
  background-color: #fff; /* 根据需要设置背景色 */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04); /* 添加阴影效果，可选 */
}

.fixed-navdiv {
  height: 53px;
}

.bluePatten {
  background-color: #2f84db;
  height: 60px;
  background-size: contain;
  background-position: center; /* 图片居中显示 */
  background: radial-gradient(717px circle at 29.71% 39.55%, #024cec 0%, #f88538 100%)
}

.bluePatten .span1 {
  font-weight: bolder;
  color: white;
  margin-top: 10px;
  padding-top: 20px;
  font-size: 30px;
  margin-left: 30px
}

.bluePatten .span2 {
  color: white;
  font-size: 18px
}

.bluePatten .span3 {
  background-color: #f44444;
  color: white;
  margin-left: 10px;
  font-size: 14px;
  padding: 4px;
  width: 80px;
  display: inline-block;
  text-align: center;
  border-radius: 16px

}

.bluePatten .span4 {
  display: block;
  margin-left: 30px;
  font-size: 13px;
  margin-top: 3px;
  color: #d1e4f7;
}
.saleType {
  display: inline-block;
  margin-left: 20px;
  margin-top: 8px;
  font-size: 10px;
  padding-bottom: 5px;
  padding: 3px;
  text-align: center;
  border-radius: 20px;
  height: 12px;
  width: 60px;
  color: white;
  letter-spacing: 2.6px;
  font-family: 'Times New Roman';
}

.carname {
  margin-top: 5px;
  margin-left: 25px;
  letter-spacing: 1.6px;
  font-family: Times New Roman;
  font-size: 18px;
  margin-bottom: 10px;
}
.MonthlyPay{
  background-color:#fff0e5;
  height:40px;
  width:360px;
  margin-left:20px;
  border-radius: 10px;
  display: flex; /* 启用Flexbox布局 */
  align-items: center; /* 垂直居中 */
}
.MonthlyPay img{
  background:radial-gradient(680px circle at 32.98% 48.31%, #ccc3bb 0%, #aba7d6 100%);
  height:40px;
  width:247px;
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
}
.MonthlyPay  span{
  align-items: center;
  font-size: 15px;
  color: orange;
  margin-left:5px;
}
.tupian1 {
  width: 378px;
  margin: 10px;
  margin-top: 20px;
  margin-bottom: 0;
}
.tupian2 {
  width: 378px;
  margin: 10px;
  margin-top: 20px;
  margin-bottom: 0;
  border-radius: 10px;
}

.img-swipe .van-swipe-item > img {
  width: 399px;
  height: auto;
  font-size: 20px;
  line-height: 150px;
  background-color: #39a9ed;
}

.list_1 li {
  width: 200px;
  margin-top: 5px;
  margin-left: 30px;
  color: #4f4f4c;
  height: 25px;
}

.list_ul {
  width: 180px;
}

.FatherSpan {
  font-size: 12px;
  width: 170px;
  margin-top: 0;
  height: 20px;
}

.childrenSpan {
  color: black;
  font-weight: bolder;
  font-size: 12px;
}

.basicInformation {
  display: flex;
  font-size: 30px;
}

img {
  width: 400px
}
.tag{
  background-color: #fff;color: #000;font-size:13px;
  margin-top: 10px;

}
</style>

