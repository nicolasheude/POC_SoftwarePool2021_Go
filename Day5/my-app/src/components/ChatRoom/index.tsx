import React from "react";

function ChatRoom(): JSX.Element {
    return (
        <div>
            <div className="all">
                <div className="allImage">
                    <img className="image"
                         src="https://media-cdn.tripadvisor.com/media/photo-s/08/60/2d/80/le-palace-de-menthon.jpg"
                         alt="Photo de profile"
                    />
                </div>
                <div className="allMessage">
                    <h4 className="name">POC FRANCE :</h4>
                    <p className="message">Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.</p>
                </div>
            </div>
        </div>
    )
}
export default ChatRoom