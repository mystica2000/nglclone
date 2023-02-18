import { createEffect, createSignal } from "solid-js"
import "../styles/Responses.css"
import DeleteIcon from "../assets/icons/delete";
import Items from "./Items";
import store from "../store/store";
import { sortHelper } from "../utils/utils";
import DeleteForResponseDialog from "../dialogs/DeleteForResponseDialog";
import { getResponses } from "../helper/getResponses";

export default function Responses() {

  const { responses, setResponses } = store;

  const [selectedOption, setSelectedOption] = createSignal('new');

  const [showDeleteAll, setShowDeleteAll] = createSignal(false);
  const [showDeleteReplied, setShowDeleteReplied] = createSignal(false);

  const getResponsesHelper = async () => {
    const value: any = await getResponses().catch(err => console.log(err))
    setResponses((value == null) ? [] : [...value]);
    sortHelper(selectedOption())
  }

  createEffect(async () => {
    let interval;
    try {
      await getResponsesHelper();

    // poll the server every 5 sec to get the new responses!
     interval = setInterval(async () => {
      await getResponsesHelper();
    }, 5000);
    } catch(e) {
     clearInterval(interval);
    }
  }, [])

  const handleSort = (e: any) => {
    setSelectedOption(e.target.value)
    sortHelper(e.target.value)
  }

  const handleDeleteAll = () => {
    setShowDeleteAll(true);
  }

  const handleDeleteReplied = () => {
    setShowDeleteReplied(true);
  }

  return (

    <div class="responses-container">

      <div style={"display:flex;flex-direction:row;margin: 20px;justify-content: flex-end;gap:1em;text-align:right;"}>
        <button class="btn btn-danger" onClick={handleDeleteAll}>Delete All Responses <DeleteIcon /> </button>
        <button class="btn btn-danger" onClick={handleDeleteReplied}>Delete Replied Responses <DeleteIcon /> </button>
      </div>
      <div class="item" style={"display: flex;flex-direction: row;justify-content: space-between;"}>

        <span>Total Responses: {responses.length}</span>
        <span>
          <label for="sort">
            sort by: &nbsp;
            <select value={selectedOption()} name="sort" id="sort" onChange={handleSort}>
              <option value="new">New</option>
              <option value="replied">Replied</option>
            </select>
          </label>
        </span>
      </div>

      {showDeleteAll() && <DeleteForResponseDialog title="Do you want to delete all the response?" okText="Yes, Delete" setShowDialog={setShowDeleteAll} id="all" dontshow={false} />}
      {showDeleteReplied() && <DeleteForResponseDialog title="Do you want to delete all the replied response?" okText="Yes, Delete"  setShowDialog={setShowDeleteReplied} id="replied" dontshow={false} />}

      <Items sortOption={selectedOption()} />

    </div>
  )
}