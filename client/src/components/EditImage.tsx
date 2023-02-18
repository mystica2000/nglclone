import { createEffect, createSignal, JSX, Show } from "solid-js";
import CloseIcon from "../assets/icons/close";
import EditIcon from "../assets/icons/edit"
import DeleteForProfileDialog from "../dialogs/DeleteForProfileDialog";
import DeleteForResponseDialog from "../dialogs/DeleteForResponseDialog";
import { updateUser } from "../helper/postUserData";
import store from "../store/store";
import "../styles/Home.css"
import { defaultProfilePicURL, updateURL } from "../utils/constants";

export default function EditImage() {

  const { user } = store

  const [pic, setPic] = createSignal("");
  const [showRemove,setShowRemove] = createSignal(false);
  const [loading,setLoading] = createSignal(false);

  createEffect(()=> {
    setPic(user()?.image || defaultProfilePicURL);
  },[user().image])

  let prof: any;
  let fileUploader: any;

  const handleProfileChange: JSX.EventHandlerUnion<HTMLInputElement, Event> = async (e: any) => {
    const image: File = e?.target?.files[0];

    // check if valid format
    if (image?.type == "image/png" || image?.type == "image/jpg" || image?.type == "image/jpeg") {
      // create a url that refers to a file so that we can use it in img!
      setLoading(true);
      let prev = pic();
      setPic(updateURL);
      try {
      const updated = await updateUser("", e.target.files[0]) // updates the user backend with current image
      console.log("test")
      if(updated) {

        setPic(URL.createObjectURL(e.target.files[0]))
        // free the memory after loading the image
        prof.onload = () => {
          // revoke to free memory
          URL.revokeObjectURL(pic())
        }
      }
      } catch(e) {
          setPic(prev);
      }
      setLoading(false);

    } else {
      window.alert("This file format not supported. \n Accepted Formats are *.png,*.jpg")
    }
  }

  const handleClick = () => {
    fileUploader.click();
  }

  const handleRemoveImage = () => {
    setShowRemove(true);
  }

  return (
    <div class="pic-container">
      <div style={"display:flex;align-items:flex-end;justify-content: end;"}><button style={"width:min-content;border:transparent;background:white"} onclick={handleRemoveImage}> <CloseIcon /></button></div>
      <label for="image">
        <img src={pic()} alt="profile picture" class="profile-pic" ref={prof} onClick={handleClick} />
        <EditIcon class="edit-icon" />
        {loading() && <div style={"position: absolute;top: 50%;left: 38%;"}>updating...</div>}
        <input style={"display: none;"} aria-label="Change Profile Picture" type="file" ref={fileUploader} accept="image/*" id="image" onChange={handleProfileChange} />
      </label>
      <div><button onClick={handleClick} style={"width:fit-content;"}>Edit Image</button></div>
      { showRemove() && <DeleteForResponseDialog title="Do you want to remove the profile picture?" okText="Yes, Remove"  setShowDialog={setShowRemove} id="" dontshow={false}  />}
    </div>
  )
}