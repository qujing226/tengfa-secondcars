<template>
  <div id="imgBox">
    <img src="https://m.tf2sc.cn/assets/img/zaixianpinggu.ae0881fd.png" alt="">
  </div>
  <div id="inputBox">
    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
            v-model="name"
            type="text"
            name="姓名"
            label="姓名"
            placeholder="姓名"
            :rules="[{ required: true, message: '请输入姓名' }]"
        />
        <van-field
            v-model="phone"

            name="手机号"
            label="手机号"
            placeholder="手机号"
            :rules="[{ required: true, message: '请输入手机号' }]"
        />
        <van-field
            v-model="carModel"
            type="text"
            name="车型"
            label="车型"
            placeholder="评估车型"
            :rules="[{ required: true, message: '请输入车型' }]"
        />
        <van-field
            v-model="kiloMeters"

            name="公里数"
            label="公里数"
            placeholder="公里数"
            :rules="[{ required: true, message: '请填写公里数' }]"
        />
        <van-field
            v-model="result"
            is-link
            readonly
            name="datePicker"
            label="时间选择"
            placeholder="上牌时间"
            @click="showPicker = true"
        />
        <van-popup v-model:show="showPicker" position="bottom">
          <van-date-picker @confirm="onConfirm" @cancel="showPicker = false" />
        </van-popup>

      </van-cell-group>

      <div >
        <van-button round block type="primary" native-type="submit"  >
          立刻估值
        </van-button>
      </div>
    </van-form>
    <van-popup
        v-model:show="showBottom"
        round
        :style="{ padding: '64px' }"
    >
      <div class="dialog-content">
        预估金额{{ this.predictPrice }}万
      </div>
    </van-popup>
  </div>
  <div id="foot">
    <div class="chinese">选择在腾发卖车的好处</div>
    <img class="carPic" src="https://m.tf2sc.cn/assets/img/sellCar_good.17aac9b7.png" alt="">
    <div class="footCard">
      <div class="cardBox">
        <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEgAAAAyCAMAAADIim3SAAACWFBMVEUAAAAAAP////////+qqqqqqv+ZmcyqqtW2ttufn7+qqsazs8yiotGiudGqqr+dscSxscSqqsylpcOqqsampr+eqsKqqsKqqsqtrcKtrcynp8Shqr2qqsakrcitrcihqsOtrcWqqsmlrcOor8WqqsSmrMSorsOqqsWnrMGlqsOlqsKnp8OnrMOorcSnp8Gnq8Gnq8akqcKnq8SmqsKmqsGpqcKmqsOoq8SlrMSmqcKnqsKmqcSmrMSoq8KmqcOmqcKmqcSnqsKmq8OnqcOnrMKnrMSlqsKoq8Omq8KnrMOmq8OnqcGoqsSnq8OmqsGmqsOmq8KnqcKnq8Knq8OmqsKnq8OlqcOmqsOnq8KmqsOnq8OmqcKmq8Ooq8OnqsKnq8Omq8OnqsGnqsOnq8Omq8Koq8KnqsOmqcKlqsKmqcKmqcGmq8KmqcKmqcKmqcKmqsKnqsOmqsGlqsKnqsOnq8KmqsKmqcKmqsKnqsOmqsKmqcKmqsKnq8KmqsGmqsKnq8KmqsKmqcGmqsKmqsKmqsGmqsOmqsKnqcKmq8KmqsKnqsOmqsKnqsKmqsGmqsKmqsOlqsKmqsGmqsOmqsKnqsKmqsKmqsOmq8OmqsOmqsKmq8KlqsKnqsKmqsKnqsOmqcKmqsKmqsGnqsOmqsKmqsKmq8KmqsKnqsKmqcGmqsOmqsKmqsGmqsKmqsKnqsOmqsKmqsKmqsKmqsKmqsKmqcGmqsKlqsKnqsOnq8OmqsOmqsKmqsGnq8KmqsKmqsOmqsKmqsOmqsGmqsKmqsKmqsOmqsKnq8OmqsJzQgnAAAAAx3RSTlMAAQECAwMFBgcICQoLCwwNDQ8REhQVFRgZGRobGxwcHh8hIiMnKy8wMTM2Nzc4Ojo6Oz0/QkdISUpTVFZWWFlcX2BhYmhoaWptbnN0dXd4eHl6enx+f4OEhYeIiYyMk5SYmZmam5ucnp+goaOkp6ysra6ws7O0tbW2t7i4ubq7vL2+v8HCwsPExcbHycrLzMzNz8/Q0NLT1Nja3d7e3+Hi4uTk5ebm5+fo6Ors7O3t7u/w8fL09PX19fb3+Pj5+fr6/Pz9/f7+h9TAygAAAmZJREFUSMft1ulXVGUcwPEniEjKItL2NNtzKVPLwrLFtIXM9kwqjRZSKmkV27QwklQqoj3TUlsgsqyUTCwn6PNv9QIGZu6dmXMeTi95Xt3ne3/zOXPm3pk7IYyv8fX/r3Prmj/ZvXfH1qcWnVJqbEl7e/sjJc5f0PKP7PprzXnFB9/FxuKn7zgsd+2/udjg6ZmS0ErJ9XCRyfuVgpZlX/7jRz8MDh8uLTh52k+loDkZsLPu5BBC1XWdIHN5gcnjNioBHd0F3q4Z3pc3go8rkoNHzf9QKegG8N6xI6GsGdTmDk24+MbHv8l+AkWgD+Dvi3LKtH5oGw0NewZzr8SGgs6loDGvPQEuHNm35l/SwtBDYHZeuwLcFQd9Cp/lt8p98GYUdA54NFFfgz+rsttZV2VXd1HoanB9oj4ILkmPf1UUegCcn6gLwDUx0PMwUJmo08HdMVAHfJusZ4EnY6DvYUvqa5OBtRFQ+QC8lOq/wvoI6HjwbKp3Q2sENAk0pfoO6IiAzgaPpfr2WOhUsCrVv4NNEVA1WJ3qv8ArMVftX3gm1Q/BczH30R+wrrDfEAN9De8XvrPvjIHegl3Jehm4NgZqgIETEnUhmBkD3QZmJeo90F8VA80DtyTqq7A5xECVv8HL+fGY32FFFBRehL6aAm9zbhy0GNyb11rgi7I46MSfoWdi7g/tINwX4qBQD54eDZO/hL4zYqEzD4Dl2f3UbaA+xELh9qHnZ9uMEEKoubUXfF4dD1W8Mfwo7nnnha6DQ4eHZod4KNR0Jv9CHqkNY4HCSa/nO71XhrFBoeymPaPMQNOUEAH9BzxyZu1s+hP4AAAAAElFTkSuQmCC" alt="">
        <div class="cardBoxmax">
          <div class="cardBoxfont"> 卖价高，最快可当天拿钱 </div>
          <div class="cardBoxfoot"> (卖价高，最快可当天完成过户，当天拿钱) </div>
        </div>

      </div>
      <div class="cardBox">
        <img src="data:image/png;base64,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" alt="">
        <div class="cardBoxmax">
          <div class="cardBoxfont">  卖车过程透明，省时省心  </div>
          <div class="cardBoxfoot">  (卖车过程透明看得到，专人办理卖车服务)  </div>
        </div>

      </div>
      <div class="cardBox">
        <img src="data:image/png;base64,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" alt="">
        <div class="cardBoxmax">
          <div class="cardBoxfont"> 实体展厅卖车信得过 </div>
          <div class="cardBoxfoot">  (20年实体经营，交易有保障)  </div>
        </div>

      </div>
    </div>
  </div>
  <div id="bottom">
    <div class="ques">常见问题</div>
    <div class="bottomBox">
      <div class="bottomTitle"> 卖车需要准备哪些资料？ </div>
      <div class="bottomDetail">
        需要先在网站提供您的联系电话，车辆品牌、型号等基本信息。待服务人员联系您后，在约定时间内准备好
        <br>
        1、身份证； 6、车示标（环保标、检字标、交强险标）；
        <br>
        2、行驶证； 7、交强险单；
        <br>
        3、车辆登记证； 8、购置税本及购置税发票；
        <br>
        4、车辆钥匙； 9、购车发票/最近一次过户发票；
        <br>
        5、说明书保养手册
      </div>
    </div>
    <div class="bottomBox">
      <div class="bottomTitle"> 卖车具体的流程是怎么样的呢？ </div>
      <div class="bottomDetail">
        1、提交资料及联系方式。
        <br>
        2、预约开车到店检测车辆或直接来店检测车辆。
        <br>
        3、现金交易结款。
        <br>
        4、办理过户手续。
      </div>
    </div>
    <div class="bottomBox">
      <div class="bottomTitle">通过腾发出售车辆需要什么费用？</div>
      <div class="bottomDetail">
        在腾发寄售车辆，卖车快，卖价高。整个服务过程透明。评估师给参考售价，明确收费内容，同意后签订寄售协议。卖出的车辆同样也享受腾发的质保服务。
      </div>
    </div>
  </div>

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
</template>

<style scoped>
/* p {
 background-color: aquamarine;
 width:195px;
 height: 200px;
} */
#imgBox img{
  width: 400px;
}
#inputBox{
  position: absolute;
  top: 100px;
  left:15px;
  width:360px;
}
.van-popup van-popup--bottom{
  width: 390px;
}
/* #app{
  height: auto;
  width: 100vw;
} */
#foot {
  position:absolute;
  top: 400px;
  left: 0px;
}
#foot .chinese{
  position: relative;
  left:80px;
  width: 200px;
  font-size: 20px;
}
#foot .carPic{
  position: relative;
  left:90px;
  width: 220px;
}
#foot .footCard{
  position: relative;
  left: 18px;
  background-color: #f5f6fa;
  width: 91vw;
  height: 51vw;
}
#foot .cardBox{
  width: auto;
  height: 65px;
  display: flex;
}
.footCard img{
  position: relative;
  top:20px;
  left: 20px;
  width: 30px;
  height: 25px;
}
.cardBoxmax{
  position: relative;
  top: 15px;
  left:40px;
  height:60px;
}
.cardBox .cardBoxfont1{
  position: relative;
  left: 50px;
  width: 300px;
  height: 20px;
  font-size: 3.8vw;
  font-weight: 400;
  color: #333;
}
#bottom {
  position: absolute;
  top: 800px;
  width: 100%;
}
#bottom .ques{
  font-weight: 600;
  font-size: 20px;
  margin-bottom: 6vw;
}
#bottom .bottomBox{
  background-image: url(https://m.tf2sc.cn/assets/img/sellCar_bg1.c8d40d6e.png);
  background-size: cover;
  width: 100%;
  height: 215px;
}
#bottom .bottomTitle{
  height: 35px;
  margin-left: 10px;
  position: relative;
  top: 20px;
  font-size: 20px;
}
#bottom .bottomDetail{
  width: 290px;
  font-size: 13px;
  position: relative;
  top:25px;
  left:15px
}
/* 价格弹窗 */
.dialog-content{
  text-align: center;
  position: relative;
  top:0px;
  font-size: 25px;
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
.custom-select > .el-select__wrapper {
  height: 200px;
}

 .custom-select > .el-select__wrapper {
   height: 30px;
   font-size: 14px;
 }

.custom-select > .el-select__input {
  height: 16px;
}
</style>
<script>
export default {
  data(){
    return{
      name:'',
      phone:"",
      kiloMeters:"",
      carModel:"",
      showPicker:false,
      result:"",
      predictPrice:"",
      showBottom:false,
      tabbar:[
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
      tabbarIndex: 3,
    }
  },

  methods:{
    TipToChangPage(ind){
      if (ind==0){
        this.$router.push(`/`)
      }else if (ind==2){
        this.$router.push(`/NewCar/`)
      }
    },
    // getPrice(getParam){
    //   // 等接口
    //   axios.get("127.0.0.1:8080/SendPredictPrice",{})
    //   .then((res)=>{
    //   predictPrice=res.data.predictPrice;
    //  this.showBottom=true,
    // })
    // },
    test23(getParam){
      this.predictPrice=25.7;
      this.showBottom=true;
    },
    onConfirm({ selectedValues }){
      this.result = selectedValues.join('/');
      this.showPicker = false;
    },
    onSubmit(values){
      // values存所有输入框值
      console.log('submit',values);
      // 记得把注释拿掉
      // getPrice(result);
      // 测试函数，拿到接口注释掉
      this.test23(values);
    }
  }

}
</script>