import React, { Component } from 'react';
// import logo from './logo.svg';
import './Submission.css';
import { Navbar, Nav, Form, FormControl, Button, Card, ListGroupItem, ListGroup, Container, Row, Col, CardGroup} from 'react-bootstrap';
import {withRouter} from 'react-router-dom';
class Submission extends Component {



    constructor(props) {
        super(props);
        this.state = {
    
        }
      }

render(){
    return (
    
        <div>
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


            <div class = "outer">
                <div class = "inner">
                <Container >
                    <Row>
                    <Col> 
                        <Card className="text-center" body>
                        {/* <Card.Img variant="top" src="holder.js/100px180?text=Image cap" /> */}
                            <Card.Body>
                                {/* <Card.Title  as="h1"> Apple </Card.Title> */}
                                <Card.Text>
                                <Card.Img variant="top" src="holder.js/100px180?text=Image cap" />
                                For contest IMAGE
                                </Card.Text>

                                <Container as="h3">
                                    <Row>
                                    <Col>Upload Photo</Col>
                                    <Col>
                                    <Button variant="outline-primary">Select File</Button>
                                    </Col>
                                    </Row>
                                </Container>
                                </Card.Body>
                            </Card>
                     </Col>
                     <Col>
                        <Card className="text-left" body>
                        {/* <Card.Img variant="top" src="holder.js/100px180?text=Image cap" /> */}
                        <Card.Body>
                            {/* <Card.Title  as="h1"> Apple </Card.Title> */}
                            <Card.Text>
                            {/* <Card.Img variant="top" src="holder.js/100px180?text=Image cap" /> */}
                            Add it to every page that displays images. I also recommend you include on your About page because, as the legal benefits reflect, the notice actually provides real information to the visitor and the About page is a logical place to give copyright information about the work on the website.

    According to the Law the copyright notice should be affixed in such a way as to “give reasonable notice of the claim of copyright.” The three elements of the notice should ordinarily appear together on the copies.

    For more information, see the Copyright Office’s Circular 3, Copyright Notice (pdf).

    You will need to change the notice every time either the earliest or latest dates of publication change. Typically you will need to change at least the second date every year if you are adding new images.
                            </Card.Text>

                            <Container as="h3">
                                <Row>
                                <Col xs={6}>Release of CopyRight</Col>
                                <Col></Col>
                                <Col>
                                    <Form.Group controlId="formBasicChecbox">
                                    <Form.Check type="checkbox" label="I agree" />
                                    </Form.Group>
                                </Col>
                                </Row>
                            </Container>
                            </Card.Body>
                        </Card>
                    </Col>
                    </Row>

                    <Row>
                        <Col></Col>

                            <Col xs={20}>
                                <div className='mt-5'>
                                <Button className="mx-auto" variant="primary" size="lg" > Submit Photo</Button> 
                                </div>
                            </Col>
                        <Col></Col>
                    </Row>
                </Container>
            </div>
            </div>
           </div>
      )

    }
}

export default Submission;
