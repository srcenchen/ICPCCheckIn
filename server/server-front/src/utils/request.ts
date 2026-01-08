import axios from "axios";
import type { Device } from "@/model/Device";

interface Response {
    devices: Device[];
}

class Requset {
    private static baseUrl: string = "";

    static async getAllDevices(): Promise<Response> {
        const res = await axios.get(`${this.baseUrl}/v1/device/all-devices`);
        return res.data;
    }

    static async deleteDeviceByMac(mac: string) {
        const res = await axios.post(`${this.baseUrl}/v1/device/delete`, {
            mac
        });
        return res.data;
    }

    static async deleteAllDevices() {
        const res = await axios.post(`${this.baseUrl}/v1/device/delete`, {
            type: "all"
        });
        return res.data;
    }

    static async getExportUrl() {
        const res = await axios.get(`${this.baseUrl}/v1/export`);
        return res.data;
    }
}

export default Requset;