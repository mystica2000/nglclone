import "../styles/Responses.css"
import "../styles/Home.css"

export function ImageDiv(props: { text: string, img: string, name: string }) {
  return (
    <div class="image-div">
      <div class="item-image">
        <p style={"min-height: 100px;margin:10px;padding:5px;display: flex;justify-content: center;align-items: center;text-align: center;height: 100%;"}>
          {props.text}
        </p>
      </div>
      <div class="user">
        <div class="header">
          <img src="https://res.cloudinary.com/dtr0c0lk2/image/upload/v1675094771/rjtzjoit2fhecgmlk8xd.jpg" alt="profile picture" class="small-img" />
          <h2>{props.name}</h2>
        </div>
      </div>
    </div>
  )
}