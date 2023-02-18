export interface User {
  name:string;
  email:string;
  active:number;
  image:string;
}

export interface UserResponse {
  id: number;
  done: number;
  response: string;
}