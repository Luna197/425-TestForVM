import React, { Component } from 'react';
// import logo from './logo.svg';
import './Contest.css';
import { Navbar, Nav, Form, FormControl, Button, Card, ListGroupItem, ListGroup, Container, Row, Col, CardGroup} from 'react-bootstrap';
import {withRouter, Redirect, Route } from 'react-router-dom';



class Contest extends Component {

  constructor(props) {
    super(props);
    this.state = {

    }
  }



handleSubmit = () => 
  { this.props.history.push('/submission')}

  render() {
    return (
    <div >
    <Navbar bg="dark" variant="dark">
        <Navbar.Brand href="#home">
        <img 
        src="./logo.svg"
        width="30"
        height="30"
        className="d-inline-block align-top"
        // alt="logo"
        />
        </Navbar.Brand>
        <Nav className="mr-auto">
          <Nav.Link href="#home">Contest</Nav.Link>
          <Nav.Link href="#features">My Submission</Nav.Link>
        </Nav>
        <Form inline>
          <FormControl type="text" placeholder="Search" className="mr-sm-2" />
          <Button variant="outline-info">Search</Button>
          <Nav.Link href="#pricing">Log Out</Nav.Link>
        </Form>
      </Navbar>

      <Container>
      <Row>
        <Col></Col>
        <Col xs={9}>
        
        
      <Card className="text-center">
      {/* <Card.Img variant="top" src="holder.js/100px180?text=Image cap" /> */}
      <Card.Body>
        <Card.Title  as="h1"> Apple </Card.Title>
        <Card.Text>
        <Card.Img variant="top" src="holder.js/100px180?text=Image cap" />
          For contest IMAGE
        </Card.Text>

          <Container as="h2">
            <Row>
              <Col>Big Cities</Col>
              <Col xs={5}></Col>
              <Col>$750</Col>
            </Row>
        </Container>

      </Card.Body>
        <ListGroup className="list-group-flush" >
          <ListGroupItem className="text-left">The big city is used to refer to a large city which seems attractive to someone because they think there are many exciting things to do there, and many opportunities to earn a lot of money.</ListGroupItem>

          <Container as="h4">
            <Card.Title className="text-left font-weight-bold" >  Requirement </Card.Title>
        </Container>

          <CardGroup>
          <Card>
            <Card.Body>
            <Card.Text> 5 </Card.Text>
              <Card.Title className="font-weight-bold" >Submission</Card.Title>
            </Card.Body>
          </Card>
          <Card>
            <Card.Body>
            <Card.Text>
                iPhone 7 or above
                </Card.Text>
              <Card.Title className="font-weight-bold">device</Card.Title>
            </Card.Body>
          </Card>
          <Card>
            <Card.Body>

        <Card.Text> 4+ </Card.Text>
          <Card.Title className="font-weight-bold">Badge Level</Card.Title>
            </Card.Body>
          </Card>
          </CardGroup>

        <Container as="h4">
            <Card.Title className="text-left font-weight-bold" >  CopyRight </Card.Title>
        </Container>

          <CardGroup>
          <Card>
            <Card.Body>
            <Card.Text> No </Card.Text>
              <Card.Title className="font-weight-bold">Ownership</Card.Title>
            </Card.Body>
          </Card>
          <Card>
            <Card.Body>
            <Card.Text> Advertisement </Card.Text>
              <Card.Title className="font-weight-bold" >Use</Card.Title>
            </Card.Body>
          </Card>
          <Card>
            <Card.Body>
              <Card.Text> 1 year </Card.Text>
              <Card.Title className="font-weight-bold">Time</Card.Title>
            </Card.Body>
          </Card>
          </CardGroup>


        <Container as="h4">
            <Card.Title className="text-left font-weight-bold" >  Similar Contests </Card.Title>
        </Container>

          <CardGroup>
          <Card>
            <Card.Img variant="top" src="holder.js/100px180?text=Image cap" />
            <Card.Body>
            <Card.Text> 
              
            <Container>
                    <Row>
                    <Col >$500</Col>
                    <Col className="font-italic"> 22 hours left</Col>
                    </Row>
                </Container>

              </Card.Text>
              <Card.Title className="font-weight-bold">Golden Gate</Card.Title>
            </Card.Body>
          </Card>

          <Card>
            <Card.Img variant="top" src="holder.js/100px180?text=Image cap" />
            <Card.Body>
            <Card.Text> 
              
            <Container>
                    <Row>
                    <Col >$500</Col>
                    <Col className="font-italic"> 22 hours left</Col>
                    </Row>
                </Container>
              
             </Card.Text>
              <Card.Title className="font-weight-bold">Golden Gate</Card.Title>
            </Card.Body>
          </Card>

          <Card>
            <Card.Img variant="top" src="holder.js/100px180?text=Image cap" />
            <Card.Body>
            <Card.Text> 

                <Container>
                    <Row>
                    <Col >$500</Col>
                    <Col className="font-italic"> 22 hours left</Col>
                    </Row>
                </Container>
  
              </Card.Text>
              <Card.Title className="font-weight-bold">Golden Gate</Card.Title>
            </Card.Body>
          </Card>
          </CardGroup>


          <Card>
        
        <Card.Body>
        {/* {this.renderRedirect()} */}
        <Button variant="primary" size="lg" onClick={this.handleSubmit}> Submission </Button> 
        </Card.Body>
          </Card>

        </ListGroup>
    </Card>
        </Col>
        <Col></Col>
      </Row>
    </Container>
    </div>
    )
  }
}

export default Contest;
