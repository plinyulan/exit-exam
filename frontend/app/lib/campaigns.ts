import { request } from "./index";

export function getCampaigns() {
  return request("/campaigns");
}
