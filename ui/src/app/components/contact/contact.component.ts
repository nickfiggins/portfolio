import { Component, Input, OnInit } from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';
import {EmailService} from './email.service';
import { ToastRef, ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.scss']
})
export class ContactComponent implements OnInit {

  @Input() type = 'general';
  shouldShowTutorQuestions : boolean = true;

  email = new FormControl('', [Validators.required, Validators.email]);
  firstName = new FormControl('', [Validators.required]);
  lastName = new FormControl('', [Validators.required]);
  bodyText = new FormControl('', [Validators.required]);

  formGroup = new FormGroup({email: this.email, 
    firstName: this.firstName,
    lastName: this.lastName,
    bodyText: this.bodyText
    })
  constructor(public emailService: EmailService,
    private toastr: ToastrService) { }


  ngOnInit(): void {
    this.shouldShowTutorQuestions = (this.type === 'tutoring') ? true : false;
    console.log(this.type);
  }

  getErrorMessage() {
    if (this.email.hasError('required')) {
      return 'You must enter a value';
    }

    return this.email.hasError('email') ? 'Not a valid email' : '';
  }

  submit(){
    if(this.formGroup.valid){
      let formInfo = this.formGroup.value;
      let emailVM = {
        email: formInfo.email,
        first_name: formInfo.firstName,
        last_name: formInfo.lastName,
        request_type: this.type, 
        message: formInfo.bodyText
      }
      this.emailService.sendEmailLambda(emailVM).subscribe(
        (response) => {
          this.toastr.success("Form has been submitted!");
          this.formGroup.reset();
          Object.keys(this.formGroup.controls).forEach(key => {
              let val = this.formGroup.get(key)
              if(val){
                val.setErrors(null)
              }
          });
        },
        (msg) => {
          console.log("Error sending form", msg);
          this.toastr.error("Error sending form. Please email me directly at nickfigginstutoring@gmail.com");
        }
      );
    }
  }

}
