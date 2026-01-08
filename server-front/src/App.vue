<template>
  <div class="container">
    <h1>签到管理</h1>
    <div class="actions">
      <button class="btn default" @click="exportData()">导出</button>
      <button class="btn default" @click="getAllDevices()">刷新</button>
      <button class="btn delete" @click="deleteAllDevices()">删除全部数据</button>
    </div>
    <table class="device-table">
      <tr>
        <th>Id</th>
        <th>签到时间</th>
        <th>签退时间</th>
        <th>Ip地址</th>
        <th>Mac地址</th>
        <th>状态</th>
        <th>学生姓名</th>
        <th>学生学号</th>
        <th>操作</th>
      </tr>
      <tr v-for="device in devices" :key="device.Id">
        <td>{{ device.Id }}</td>
        <td>{{ formatDate(device.CheckIn, checkType.CheckIn, device.Status) }}</td>
        <td>{{ formatDate(device.CheckOut, checkType.CheckOut, device.Status) }}</td>
        <td>{{ device.Ip }}</td>
        <td>{{ device.Mac }}</td>
        <td>{{ formatStatus(device.Status) }}</td>
        <td>{{ device.StuName }}</td>
        <td>{{ device.StuNum }}</td>
        <td>
          <button class="btn delete" @click="handleDelete(device.Mac)">删除</button>
        </td>
      </tr>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import Requset from './utils/request';
import type { Device } from './model/Device';

const devices = ref<Device[]>([]);

const getAllDevices = async () => {
  const res = await Requset.getAllDevices();
  devices.value = res.devices;
}

enum checkType {
  CheckIn,
  CheckOut
}

const formatDate = (ISOString: string, type: checkType, status: number) => {
  if (type === checkType.CheckIn && status === 0) {
    return '未签到';
  }
  if (type === checkType.CheckOut && status === 1) {
    return '未签退';
  }
  const date = new Date(ISOString);
  return date.toLocaleString();
}

const formatStatus = (status: number) => {
  if (status === 0) {
    return '未签到';
  } else if (status === 1) {
    return '已签到';
  } else if (status === 2) {
    return '已签退';
  } else {
    return '未知状态';
  }
}

const handleDelete = (mac: string) => {
  if (confirm('确定要删除该设备吗？')) {
    Requset.deleteDeviceByMac(mac);
    getAllDevices();
    alert('删除成功');
  }
}

const deleteAllDevices = async () => {
  if (confirm('确定要删除全部数据吗？')) {
    await Requset.deleteAllDevices();
    getAllDevices();
    alert('删除成功');
  }
}

const exportData = async () => {
  const res = await Requset.getExportUrl();
  window.open('/' + res.path);
}

onMounted(() => {
  getAllDevices();
})
</script>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.actions {
  width: 100%;
  display: flex;
  flex-direction: row-reverse;
  gap: 20px;
}

.device-table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
  font-size: 14px;
}

.device-table th,
.device-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #e0e0e0;
}

.device-table th {
  background-color: #f5f5f5;
  font-weight: 600;
  color: #333;
}

.device-table tr:hover {
  background-color: #f9f9f9;
}

.device-table tr:nth-child(even) {
  background-color: #fafafa;
}

.btn {
  padding: 6px 12px;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
  transition: background-color 0.3s;
}

.btn.default {
  background-color: #1890ff;
}

.btn.default:hover {
  background-color: #40a9ff;
}

.btn.delete {
  background-color: #ff4d4f;
}

.btn.delete:hover {
  background-color: #ff7875;
}
</style>
