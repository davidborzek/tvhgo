import axios from "axios";
import { ErrorResponse, UserResponse } from "./api.types";

export class ApiError extends Error {
  constructor(public readonly code: number, message: string) {
    super(message);
  }
}

const client = axios.create({
  baseURL: "/api",
});

client.interceptors.response.use(
  (res) => res,
  (err) => {
    if (axios.isAxiosError(err) && err.response) {
      const data = err.response.data as ErrorResponse;
      return Promise.reject(new ApiError(err.response.status, data.message));
    }
    return Promise.reject(err);
  }
);

export async function login(username: string, password: string) {
  await client.post("/login", {
    username,
    password,
  });
}

export async function logout(): Promise<void> {
  await client.post("/logout");
}

export async function getUser(): Promise<UserResponse> {
  const response = await client.get<UserResponse>("/user");
  return response.data;
}

export async function fetch<T>(path: string): Promise<T> {
  const response = await client.get<T>(path);
  return response.data;
}