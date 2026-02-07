import { request } from "./index";

export function getPoliticians() {
  return request("/politicians/");
}

export function getPromisesByPolitician(id: number) {
  return request(`/politicians/${id}/promises`);
}
