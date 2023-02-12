export type User = {
  id: number,
  email: string,
  name: string,
  type: "admin" | "user",
  password: string,
}
