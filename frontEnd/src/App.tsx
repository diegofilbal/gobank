import GlobalStyle from "./styles/global";
import styled from "styled-components";
import { toast, ToastContainer } from "react-toastify";
import { useEffect, useState } from "react";
import "react-toastify/dist/ReactToastify.css";
import Form from "./components/form";
import Grid from "./components/Grid";
import { reservasType} from "./components/Grid";
import axios from "axios";

const Container = styled.div`
width: 100%;
max-width: 1000px;
margin-top: 20px;
display: flex;
flex-direction: column;
align-items: center;
gap: 10px;
`;

function App() {
  
  return (
    <div className="App">
      <ToastContainer autoClose={3000}/>
      <GlobalStyle/>
      <Container>
        <Form></Form>
        <Grid ></Grid>
      </Container>
    </div>
  );
}

export default App;
