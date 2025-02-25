package interfaces

// MailDAO é o tipo concreto que implementa a interface DAO.
type MailDAO struct {
	Destination string
	Subject     string
	Body        string
}

// Caso a interface DAO exija métodos específicos, implemente-os abaixo.
// Por exemplo, se a interface DAO for declarada assim:
// type DAO interface {
//     Send() error
// }
// então implemente o método Send para MailDAO:
//
// func (m *MailDAO) Send() error {
//     // implementação do envio de email...
//     return nil
// } 