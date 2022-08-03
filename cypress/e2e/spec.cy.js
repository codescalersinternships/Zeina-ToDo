// Cypress.on('uncaught:exception', (err, runnable) => {
//   // returning false here prevents Cypress from
//   // failing the test
//   return false
//   });
describe('The homepage', () => {
  Cypress.on('uncaught:exception', (err, runnable) => {
    return false;
  });
  beforeEach(()=>{
  cy.visit('http://localhost:5656')

 }) 


  
  it('add todo', () => {
    cy.get('#new-todo').type('task23456')
    // cy.get('#toggle-all').
    cy.get('input').type('{enter}')
    cy.contains("task23456")


})
it('delete todo', () => {
  cy.get('#delete1').click()
  // cy.contains("task23456")


})

})