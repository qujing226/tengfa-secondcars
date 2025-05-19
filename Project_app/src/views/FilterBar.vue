<template>
  <van-divider content-position="left" style="height:2px;margin:5px"
               :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>
  <div class="filters">
    <label>
      <span>排序:</span>
      <select v-model="filters.sortKey" @change="emitFilters">
        <option value="default">默认排序</option>
        <option value="latest">最新发布</option>
        <option value="priceHigh">价格最高</option>
        <option value="priceLow">价格最低</option>
      </select>
    </label>
    <label>
      <span>车型:</span>
      <select v-model="filters.brand" @change="emitFilters">
        <option value="">所有车型</option>
        <option value="Mercedes">两箱轿车</option>
        <option value="BMW">三厢轿车</option>
        <option value="BMW">SUV</option>
        <!-- 添加更多品牌 -->
      </select>
    </label>
    <label>
      <span>价格:</span>
      <select v-model="filters.priceRange" @change="emitFilters">
        <option value="">所有价格</option>
        <option value="0-10">0-10万元</option>
        <option value="10-20">10-20万元</option>
        <option value="20-30">20-30万元</option>
        <!-- 添加更多价格区间 -->
      </select>
    </label>
    <label>
      <span>车龄:</span>
      <select v-model="filters.age" @change="emitFilters">
        <option value="">所有车龄</option>
        <option value="0-1">0-1年</option>
        <option value="1-3">1-3年</option>
        <option value="3-5">3-5年</option>
        <!-- 添加更多车龄区间 -->
      </select>
    </label>
    <label>
      <span>里程:</span>
      <select v-model="filters.mileage" @change="emitFilters">
        <option value="">所有里程</option>
        <option value="0-10">0-10万公里</option>
        <option value="10-20">10-20万公里</option>
        <option value="20-30">20-30万公里</option>
        <!-- 添加更多里程区间 -->
      </select>
    </label>
  </div>
  <van-divider content-position="left" style="height:2px;margin:5px"
               :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>
</template>


<script>
export default {
  data() {
    return {
      filters: {
        sortKey: 'default',
        brand: '',
        priceRange: '',
        age: '',
        mileage: ''
      }
    };
  },
  methods: {
    emitFilters() {
      this.$emit('filter', this.filters);
    }
  }
};
</script>

<style>
.filters {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  margin-bottom: 10px;
  margin-top: 10px;

}

.filters label {
  display: flex;
  flex-direction: column;
  width:70px;
  font-size:15px;

  margin-left:7px;
}
.filters span {
  color:#1a191a;
  font-weight: bolder;
  display: flex; /* 保持不变 */
  align-items: center; /* 垂直居中 */
  justify-content: center; /* 水平居中 */
  width: 35px;
  font-size: 13px;
  margin-left: 13px;
  height: 100%; /* 确保span元素填充其父容器的高度，如果父容器有固定高度的话 */
}
.filters select {
  border-style: solid;
  border-radius: 5px;
  width: 60px;
  margin: 0 5px 5px 0; /* 简化外边距属性 */
  height: 32px; /* 增加高度以适应字体大小和垂直居中需求，按需调整 */
  line-height: 30px; /* 设置行高等于高度，实现垂直居中 */
  font-size: 10px; /* 调整字体大小以匹配设计需求 */
}
</style>
