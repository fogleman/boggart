// Code generated by go-swagger; DO NOT EDIT.

package habmin_dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HTTPGetDashboardsReader is a Reader for the HTTPGetDashboards structure.
type HTTPGetDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HTTPGetDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewHTTPGetDashboardsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewHTTPGetDashboardsDefault creates a HTTPGetDashboardsDefault with default headers values
func NewHTTPGetDashboardsDefault(code int) *HTTPGetDashboardsDefault {
	return &HTTPGetDashboardsDefault{
		_statusCode: code,
	}
}

/*HTTPGetDashboardsDefault handles this case with default header values.

successful operation
*/
type HTTPGetDashboardsDefault struct {
	_statusCode int
}

// Code gets the status code for the http get dashboards default response
func (o *HTTPGetDashboardsDefault) Code() int {
	return o._statusCode
}

func (o *HTTPGetDashboardsDefault) Error() string {
	return fmt.Sprintf("[GET /habmin/dashboards/{dashboardId}][%d] httpGetDashboards default ", o._statusCode)
}

func (o *HTTPGetDashboardsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}