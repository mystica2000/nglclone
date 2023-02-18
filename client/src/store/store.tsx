import { createRoot, createSignal } from "solid-js";
import { createStore } from "solid-js/store";
import { User, UserResponse } from "../utils/types";

function userStore() {
  const [user,setUser] = createSignal<User>({name:"",email:"",active:-1,image:""});
  const [responses,setResponses] = createStore<UserResponse[]>([]);
  const [sessionExpired,setSessionExpired] = createSignal(false);

  return {user,setUser,responses,setResponses};
}

export default createRoot(userStore)