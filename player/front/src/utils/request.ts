import axios from "axios";

class Requset {
  private static baseUrl = "";

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