import { createEffect, Show } from 'solid-js';
import Navbar from '../components/Navbar';
import { nav } from '../utils/constants';
import "../styles/Home.css"
import Responses from '../components/Responses';
import Profile from '../components/Profile';
import MoveToTop from '../components/MoveToTop';
import store from '../store/store';
import { getUserData } from '../helper/getUserData';
import { Footer } from '../components/Footer';

export default function Home() {

  const {user} = store

  createEffect(async ()=> {
    await getUserData();
  },[])

  return (<Show
      when={user().name.length>0}
      fallback={<div>display like session logged out, try again!</div>}
    >
      <Navbar displayFor={nav.HOME} />
      <div class="homepage">
      <div class="home-container">
        <Profile />
        <Responses />
        <MoveToTop />
      </div>
      </div>
      <Footer />
    </Show>)
}