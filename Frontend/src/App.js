import { BrowserRouter as Router,Routes,Route } from 'react-router-dom'
import { Users } from './Container/Users/ViewUsers'
import { Document } from './Container/Documents/ViewDocument/index.js'
import { AddDocument } from './Container/Documents/CreateDocument'
import { Createuser } from './Container/Users/CreateUser'
import { EditUsers } from './Container/Users/ViewUsers/edit'
import { EditDocuments } from './Container/Documents/ViewDocument/edit'

 function App(){
  return(
      <>
      <Router>
          <Routes>
              <Route path='/' exact element={<Users/>}/>
              <Route path='/createuser' element={<Createuser/>}/> 
              <Route path='/documents' element={<Document/>}/>
              <Route path='/insertdoc' element={<AddDocument/>}/>
              <Route path='/editusers' element={<EditUsers/>}/>
              <Route path='/editdocuments' element={<EditDocuments/>}/>
          </Routes>
      </Router>
      </>
  )
}
export default App