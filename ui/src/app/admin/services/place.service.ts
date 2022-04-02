import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http'
import 'rxjs/add/operator/catch';
import { retry, catchError } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class PlaceService {
  places: any = [];

  constructor(private http: HttpClient) { }

 // Http Options
 httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
}

// Handle API errors
handleError(error: HttpErrorResponse) {
  console.log('?')
  if (error.error instanceof ErrorEvent) {
    // A client-side or network error occurred. Handle it accordingly.
    console.error('An error occurred:', error.error.message);
  } else {
    // The backend returned an unsuccessful response code.
    // The response body may contain clues as to what went wrong,
    console.error(
      `Backend returned code ${error.status}, ` +
      `body was: ${error.error}`);
  }
  // return an observable with a user-facing error message
  return throwError(
    'Something bad happened; please try again later.');
};

  getAllPlaces(){
    return this.http.get('/api/places');
  }


  deletePlace(id:string){
    return this.http.delete("/api/places" + id).subscribe(
      (res) => {
        return res;
      }
    )
  }

  addPlace(name:string, folderId: string){
    this.http.post('/api/places', {
      "name": name,
      "folder_id": folderId
    }, this.httpOptions).subscribe(
      (res) => {
        return res;
      }
    )
  }
}
