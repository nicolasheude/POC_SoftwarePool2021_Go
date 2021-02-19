import React from "react";

type MessageProps = {
    imgProfileUrl: string;
    messageId: number;
    senderName: string;
    messageText: string;
}

function Message(props: MessageProps): JSX.Element {
    return (
        <div>
            <h1 className="title">Conversation :</h1>
            <div className="all">
                <div className="allImage">
                    <img className="image"
                    src={props.imgProfileUrl}
                    alt="Photo de profile"
                />
            </div>
            <div className="allMessage">
                <h4 className="name">{`${props.senderName} :`}</h4>
                <p className="message">{`${props.messageText}`}</p>
            </div>
            </div>
        </div>
    )
}
export default Message