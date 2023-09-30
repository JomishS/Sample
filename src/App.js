import { BrowserRouter as Router,Routes,Route } from 'react-router-dom'
import { Users } from './Container/Users'
import { Document } from './Container/Documents'
import { AddDocument } from './Container/Documents/AddDocument'
import { Createuser } from './Container/Users/Createuser'

 function App(){
  return(
      <>
      <Router>
          <Routes>
              <Route path='/' exact element={<Users/>}/>
              <Route path='/createuser' element={<Createuser/>}/> 
              <Route path='/documents' element={<Document/>}/>
              <Route path='/insertdoc' element={<AddDocument/>}/>
          </Routes>
      </Router>
      </>
  )
}
export default App