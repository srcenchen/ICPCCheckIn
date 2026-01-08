<template>
  <div class="app-container">
    <!-- 签到模式 -->
    <div v-if="mode === 'checkin'" class="checkin-container">
      <h2>学生签到</h2>
      <p>严禁场外答题，否则成绩无效！</p>
      <div class="form-group">
        <label>姓名：</label>
        <input v-model="stuName" type="text" placeholder="请输入姓名">
      </div>
      <div class="form-group">
        <label>学号：</label>
        <input v-model="stuNum" type="number" placeholder="请输入学号">
      </div>
      <button @click="handleCheckIn" class="btn-primary">签到</button>
    </div>

    <!-- 签退模式 -->
    <div v-else-if="mode === 'checkout'" class="checkout-container">
      <h2>学生签退</h2>
      <p>姓名：{{ stuName }}</p>
      <p>学号：{{ stuNum }}</p>
      <button @click="handleCheckOut" class="btn-danger">签退</button>
    </div>

    <!-- 已签退 -->
    <div v-else-if="mode === 'checkover'" class="loading-container">
      <p>您已签退</p>
      <p>您的电脑将在1分钟内自动关机。</p>
    </div>

    <!-- 加载状态 -->
    <div v-else class="loading-container">
      <p>正在加载签到状态...</p>
    </div>

    <!-- 错误状态 -->
    <div v-if="mode === 'error'" class="loading-container">
      <p>加载失败，请联系工作人员</p>
    </div>

    <!-- 签退成功警告 -->
    <!-- <div v-if="showWarning" class="warning-modal">
      <div class="warning-content">
        <h3>⚠️ 重要提醒 ⚠️</h3>
        <p>严禁场外答题，否则成绩无效！</p>
        <p>您的电脑将在1分钟内自动关机。</p>
        <button @click="showWarning = false" class="btn-warning">我知道了</button>
      </div>
    </div> -->
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import Requset from './utils/request';

const mode = ref<'checkin' | 'checkout' | 'checkover' | 'loading' | 'error'>('loading');
const stuName = ref('');
const stuNum = ref('');
const showWarning = ref(false);

const checkStatus = async () => {
  try {
    const status = await Requset.getStatus();
    switch (status.code) {
      case 0: mode.value = 'checkin'; break;
      case 1: {
        mode.value = 'checkout';
        stuName.value = status.stuName;
        stuNum.value = status.stuNum.toString();
        break
      };
      case 2: mode.value = 'checkover'; break;
      default: mode.value = 'error'; break;
    }
  } catch (error) {
    console.error('获取状态失败:', error);
    mode.value = 'error';
  }
};

const handleCheckIn = async () => {
  if (!stuName.value || !stuNum.value) {
    alert('请填写完整的姓名和学号');
    return;
  }

  try {
    if (confirm(`请确认您输入的信息无误：\n姓名:${stuName.value}\n学号:${stuNum.value}`)) {
      await Requset.cin(stuName.value, parseInt(stuNum.value));
      mode.value = 'checkout';
      alert('签到成功！');
    }
  } catch (error) {
    console.error('签到失败:', error);
    alert('签到失败，请重试');
  }
};

const handleCheckOut = async () => {
  try {
    if (confirm("是否确认签退？")) {
      await Requset.cout();
      showWarning.value = true;
      mode.value = 'checkover';
      stuName.value = '';
      stuNum.value = '';
    }
  } catch (error) {
    console.error('签退失败:', error);
    alert('签退失败，请重试');
  }
};

onMounted(() => {
  checkStatus();
});
</script>

<style scoped>
.app-container {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}

.checkin-container,
.checkout-container,
.loading-container {
  text-align: center;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.form-group {
  margin: 20px 0;
  text-align: left;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

.form-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

.btn-primary,
.btn-danger,
.btn-warning {
  padding: 10px 30px;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  margin-top: 20px;
}

.btn-primary {
  background-color: #4CAF50;
  color: white;
}

.btn-danger {
  background-color: #f44336;
  color: white;
}

.btn-warning {
  background-color: #ff9800;
  color: white;
}

.warning-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.warning-content {
  background-color: white;
  padding: 30px;
  border-radius: 8px;
  max-width: 400px;
  text-align: center;
}

.warning-content h3 {
  color: #f44336;
  margin-bottom: 20px;
}

.warning-content p {
  margin: 10px 0;
  line-height: 1.6;
}
</style>
