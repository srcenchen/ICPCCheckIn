import axios from "axios";

interface statusResponse {
    /**
     * 签到时间
     */
    checkInTime: string;
    /**
     * 签退时间
     */
    checkOutTime: string;
    /**
     * 0 代表未签到 1代表已经签到 2代表已经签退
     */
    code: number;
    /**
     * IP地址
     */
    ip: string;
    /**
     * MAC地址
     */
    mac: string;
    /**
     * 姓名
     */
    stuName: string;
    /**
     * 学号
     */
    stuNum: string;
    [property: string]: any;
}

interface cinResponse {
    /**
     * 0 代表成功 1代表失败
     */
    code: number;
    /**
     * 消息
     */
    message: string;
    [property: string]: any;
}

interface coutResponse {
    /**
     * 0 成功 1失败
     */
    code: number;
    message: string;
    [property: string]: any;
}



class Requset {
  private static baseUrl = "http://localhost:3000/api";

  static async getStatus() {
    const response = await axios.get(this.baseUrl + "/v1/status");
    return response.data;
  }

  static async cin(stuName: string, stuNum: number) {
    const response = await axios.post(this.baseUrl + "/v1/check-in", {
      stuName,
      stuNum: stuNum.toString(),
    });
    return response.data;
  }

  static async cout() {
    const response = await axios.post(this.baseUrl + "/v1/check-out");
    return response.data;
  }
}

export default Requset;