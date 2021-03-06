// @ts-nocheck
interface Pub {
  Quantity: number;
  Publications: {
    id: number;
    title: string;
    miniature: string;
    bodyOfDocument: string;
  }[];
  Size: number;
}

function NewPublications() {
  let urlApi: string =
    "api" + "/" + window.location.pathname.replace(/\//gi, "");
  let publication: Pub | string = {
    Publications: [
      {
        id: 0,
        title: "",
        miniature: "",
        bodyOfDocument: "",
      },
    ],
    Size: 0,

    Quantity: 0,
  };

  // this is only for get the api
  fetch(urlApi)
    .then((r) => r.text())
    .then((data) => {
      publication = data;
      publication = JSON.parse(publication);
    });

  let d: any = document.getElementById("publications");

  // add the elements
  for (let i of publication.Publications) {
    // then add elements into the dom
    let element = `
    <p>
      <a  class="publications" href="/publication/${i.id}">
        <div >
          <div class="publicationContent">
          <img src="${i.miniature}" >    
          <div class="aboutPubication"> 
            <h1 >
              ${i.title}
            </h1>
            <p >
            ${
              i.bodyOfDocument
                .slice(0, i.bodyOfDocument.indexOf("\n") % 40)
                .replace(
                  /#|'|`|"|\||-|@|=|([(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*))/gi,
                  ""
                ) + "..."
            }
            </p>
          </div>
            
            </div>
          </div>
        </div>
      </a>
    </p>
      
    `;

    d.innerHTML += element;
  }
  let pagePublications: any = document.getElementById("pagePublications");
  for (
    let i: number = 1;
    i <= publication.Size / publication.Quantity + 1;
    i++
  ) {
    let Element: string = `
 
 
      <a   href="/${i}"  >
        <div class="buttonElementID"  >
        <h3 > ${i} </h3>
        </div>
      </a>
 
    `;

    pagePublications.innerHTML += Element;
    console.log(i); // this supose to be you
  }
}
NewPublications();
