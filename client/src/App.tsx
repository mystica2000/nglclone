import { Component, lazy } from 'solid-js';
import { Routes,Route } from '@solidjs/router'
const Login:any = lazy(()=> import("./pages/Login"))
const Home:any = lazy(()=> import("./pages/Home"))
const Form:any = lazy(()=> import("./pages/Form"))
const Name:any = lazy(()=> import("./pages/Name"))
const Error:any = lazy(()=> import("./pages/Error"))
const Redirect:any = lazy(()=> import("./pages/Redirect"))
const SessionExpired = lazy(()=> import("./pages/SessionExpiredPage"))
import axios from "axios";

axios.defaults.baseURL = import.meta.env.VITE_API_URL;

const App: Component = () => {
  return (
    <>
    <Routes>
    <Route path="/" component={Login} />
    <Route path="/home" component={Home} />
    <Route path="/name" component={Name} />
    <Route path="/:id" component={Form} />
    <Route path="/redirect" component={Redirect} />
    <Route path="/error" component={Error} />
    <Route path="/expired" component={SessionExpired} />
  </Routes>
  </>
  );
};

export default App;