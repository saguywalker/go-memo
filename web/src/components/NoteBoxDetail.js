import React, { useState } from "react";
import NoteBoxEdit from "./NoteBoxEdit.js";

function NoteBoxDetail({ note, onCloseNoteBoxDetail, setEditValueCallback }) {
  const [isEdit, setIsEdit] = useState(false);
  
  return (
	<div className="Note-box-detail-wrapper">
		<div className="Note-box detail">
			{
				!isEdit ?
					<div>
						<div className="box-header">
            	<h2>{note.title}</h2>
            	<p>Last edit: {note.lastEdit}</p>
          	</div>
            {note.detail.split ('\n').map ((item, i) => <p key={i}>{item}</p>)}
						<button type="button" className="btn btn-primary" onClick={() => setIsEdit(true)}>Edit</button>
						<button type="button" className="ml-2 btn btn-danger" onClick={() => onCloseNoteBoxDetail(false)}>Close</button>
					</div> :
					<NoteBoxEdit
						note={note}
						setIsEditCallback={(isEditData) => setIsEdit(isEditData)}
						setEditValueCallback={(editValue) => setEditValueCallback(editValue)}
						closeNoteBoxDetail={() => onCloseNoteBoxDetail(false)}
					/>
			}
		</div>
	</div>
  )
}

export default NoteBoxDetail;