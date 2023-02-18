import "../styles/Responses.css"
import MarkIcon from "../assets/icons/mark"
import DeleteIcon from "../assets/icons/delete"
import { createEffect, createSignal, For, onMount, Show } from "solid-js"
import CameraIcon from "../assets/icons/camera"
import UpArrowIcon from "../assets/icons/up"
import DownArrowIcon from "../assets/icons/down"
import axios, { AxiosResponse } from "axios"
import store from "../store/store"
import { sortHelper } from "../utils/utils"
import { deleteQuestion } from "../helper/deleteQuestion"
import DeleteForResponseDialog from "../dialogs/DeleteForResponseDialog"
import html2canvas from "html2canvas"
import { defaultProfilePicURL } from "../utils/constants"
import TwitterIcon from "../assets/icons/twitter"




export default function Items(props: { sortOption: string }) {

  const { responses, setResponses } = store;
  let imageDivElement: any;
  const { user } = store

  const [showUpArrowArray, setShowUpArrowArray] = createSignal<boolean[]>(new Array(responses.length).fill(false));
  const [showDialog, setShowDialog] = createSignal<boolean>(false);
  const [showImage, setShowImage] = createSignal(false);

  const ImageAsDiv = (props:{currentResponse:string}) => {
    return (
      <div class="image-div" ref={imageDivElement} style={showImage() ? "display:block" : "display:none"} >
      <div class="item-image">
        <p style={"min-height: 100px;margin:10px;padding:5px;display: flex;justify-content: center;align-items: center;text-align: center;height: 100%;"}>
          {props.currentResponse}
        </p>
      </div>
      <div class="user">
        <div class="header">
          <div class="small-img-mod" style={`background-image:url(${user()?.image || defaultProfilePicURL})`}></div>
          <h2>{user().name}</h2>
        </div>
      </div>
    </div>
    )
  }

  const shrinkOrExpandEvent = (index: number) => {
    let value = !showUpArrowArray()[index];

    setShowUpArrowArray(() => {
      let newArray = [...showUpArrowArray()];
      newArray[index] = value;
      return newArray;
    })
  }

  const handleReply = async (index: number, id: number) => {

    axios.put<number>("/response", id, { withCredentials: true }).then(function (response: AxiosResponse) {
      // handle success
      if (response.status == 200) {
        const updateResponses = responses.map((res) => {
          if (res.id == id) {
            return { ...res, done: Number(response.data) }
          }
          return res;
        })
        setResponses(updateResponses);
        sortHelper(props.sortOption);

        setShowUpArrowArray(new Array(responses.length).fill(false));
      }
    })
      .catch(function (error: unknown) {
        // handle error
        if (axios.isAxiosError(error)) {
          if (error.response) {
            if (error.response.status == 423) {
              window.location.href = "/expired"
            }
          } else if (error.request) {

          }
        } else {
          console.log(error)
        }
      })
  }

  const handleDelete = async (id: number) => {
    // localstorage
    // call the dialog component;
    let dontshow = localStorage.getItem("dontshow") || null;
    if (dontshow == "1") {
      try {
        await deleteQuestion(id);
      } catch (e) {
        // something went wrong, try again!
      }
    } else {
      setShowDialog(true);
    }

  }

  const handleDownload = async (aRow: any) => {
    setShowImage(true);
    html2canvas(imageDivElement, {
      useCORS: true,
    }).then(function (canvas) {
      var dataURL = canvas.toDataURL("image/png");
      var downloadLink = document.createElement("a");
      downloadLink.href = dataURL;
      downloadLink.download = "myImage.png";
      downloadLink.click();
      setShowImage(false);
    });
  }

  return (
    <ul >
      <For each={responses}>
        {(res, i) => {

          return (<li class="item" classList={{ markedDone: res.done == 1 }}>

            <button style={"display:flex;flex-direction:column;align-items:center;cursor:pointer;background-color:inherit;border:none;"}
              onClick={() => { shrinkOrExpandEvent(i()) }}>
              <Show when={showUpArrowArray()[i()]}
                fallback={<DownArrowIcon />}>
                <UpArrowIcon />
              </Show>
            </button>

            <Show when={showUpArrowArray()[i()]}
              fallback={<div style={"display:grid;grid-template-columns:1fr 0fr;"}>
                <p class="pbox psmall" style={"min-height:50px;"}>{res.response}</p>
                <ImageAsDiv currentResponse={res.response}/>
              </div>
              }>
              <div style={"display:grid;grid-template-columns:1fr 0fr;"}>
                <p class="pbox" style={"min-height: 100px;"}>{res.response}</p>

              </div>

              <ImageAsDiv currentResponse={res.response}/>

              <span class="btn-header">
                <button class="btn btn-row btn-warning" onClick={() => handleReply(i(), res.id)}>{res.done == 0 ? "Mark as Replied" : "Mark as Not replied"} <MarkIcon />  </button>
                <button class="btn btn-row btn-success" onClick={() => handleDownload(res)} >Download <CameraIcon />   </button>
                <button class="btn btn-row btn-danger" onClick={() => handleDelete(res.id)} >Delete Response <DeleteIcon /> </button>
              </span>

              {showDialog() && <DeleteForResponseDialog okText="Yes, Delete" dontshow={true} title="Do you want to delete the response?" setShowDialog={setShowDialog} id={res.id} />}

            </Show>
          </li>)
        }
        }</For>
    </ul>
  )
}
